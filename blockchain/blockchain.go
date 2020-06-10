package blockchain

import (
	"fmt"
	"github.com/han-joker/bc-demo/block"
	"github.com/han-joker/bc-demo/pow"
	"github.com/han-joker/bc-demo/tx"
	"github.com/han-joker/bc-demo/wallet"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"time"
)

type BlockChain struct {
	lastHash block.Hash  // 最后一个区块的哈希
	db       *leveldb.DB // leveldb 的连接
}

// 构造方法
func NewBlockchain(db *leveldb.DB) *BlockChain {
	// 实例化 Blockchain
	bc := &BlockChain{
		db: db,
	}

	// 初始化 lastHash
	// 读取最后的区块哈希
	data, err := bc.db.Get([]byte("lastHash"), nil)
	if err == nil { // 读取到 lasthash
		bc.lastHash = block.Hash(data)
	}
	return bc
}

// 添加创世区块（第一个区块）
func (bc *BlockChain) AddGensisBlock(address wallet.Address) *BlockChain {
	// 校验是否可以添加创世区块
	if bc.lastHash != ""  {
		// 已经存在区块，不需要再添加创世区块
		return bc
	}

	// 只有 txs 是特殊
	return bc.AddBlock(address)
}

// 添加区块
// 提供区块的数据，目前是字符串
// address 添加该区块的地址
func (bc *BlockChain) AddBlock(address wallet.Address) *BlockChain {
	// # 构建区块
	b := block.NewBlock(bc.lastHash)

	// # 增加 CoinBase 交易
	cbtx := tx.NewCoinbaseTX(address)
	// 将交易加入到区块中
	b.AddTX(cbtx)

	// 处理区块交易

	// 对区块做 POW，工作量证明
	// pow 对象
	p := pow.NewPOW(b)
	// 开始证明
	nonce, hash := p.Proof()
	if hash == "" {
		log.Fatal("block Hashcash Proof Failed!")
	}
	// 为区块设置nonce和hash
	b.SetNonce(nonce).SetHashCurr(hash)


	// 将区块加入到链的存储结构中
	if bs, err := block.BlockSerialize(*b); err != nil {
		log.Fatal("block can not be serialized.")
	} else if err = bc.db.Put([]byte("b_" + b.GetHashCurr()), bs, nil); err != nil {
		log.Fatal("block can not be saved")
	}

	// 将最后的区块哈希设置为当前区块
	bc.lastHash = b.GetHashCurr()
	// 将最后的区块哈希存储到数据库中
	if err := bc.db.Put([]byte("lastHash"), []byte(b.GetHashCurr()), nil); err != nil {
		log.Fatal("lastHas can not be saved")
	}
	return bc
}

// 通过hash获取区块
func (bc *BlockChain) GetBlock(hash block.Hash) (*block.Block, error) {
	// 从数据库中读取对应的区块
	data, err := bc.db.Get([]byte("b_" + hash), nil)
	if err != nil {
		return nil, err
	}
	// 反序列化
	b, err := block.BlockUnSerialize(data)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// 迭代展示区块的方法
func (bc *BlockChain) Iterate() {
	// 构建迭代器
	bci := NewBCIterator(bc)
	for b, err := bci.Next(); err == nil; b, err = bci.Next(){
		// 做 hashcash 验证
		pow := pow.NewPOW(b)
		if !pow.Validate() {
			log.Fatalf("Block <%s> is not Valid.", b.GetHashCurr())
			continue
		}

		fmt.Println("HashCurr:", b.GetHashCurr())
		fmt.Println("TXs:", b.GetTxsString())
		fmt.Println("Time:", b.GetTime().Format(time.UnixDate))
		fmt.Println("HashPrev:", b.GetHashPrevBlock())
		fmt.Println()
	}

}

func (bc *BlockChain) Clear() {
	// 数据库中全部区块链的 key 全部删除
	bc.db.Delete([]byte("lastHash"), nil)
	// 迭代删除，全部的 b_ 的key
	iter := bc.db.NewIterator(util.BytesPrefix([]byte("b_")), nil)
	for iter.Next() {
		bc.db.Delete(iter.Key(),nil)
	}
	iter.Release()

	// 清空 bc 对象
	bc.lastHash = ""
}


