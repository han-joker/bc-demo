package blockchain

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"time"
)

type BlockChain struct {
	lastHash Hash // 最后一个区块的哈希
	db *leveldb.DB // leveldb 的连接
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
		bc.lastHash = Hash(data)
	}
	return bc
}

// 添加创世区块（第一个区块）
func (bc *BlockChain) AddGensisBlock() *BlockChain {
	// 校验是否可以添加创世区块
	if bc.lastHash != ""  {
		// 已经存在区块，不需要再添加创世区块
		return bc
	}

	// 只有 txs 是特殊
	return bc.AddBlock("The Gensis Block.")
}

// 添加区块
// 提供区块的数据，目前是字符串
func (bc *BlockChain) AddBlock(txs string) *BlockChain {
	// 构建区块
	b := NewBlock(bc.lastHash, txs)
	// 将区块加入到链的存储结构中
	if bs, err := BlockSerialize(*b); err != nil {
		log.Fatal("block can not be serialized.")
	} else if err = bc.db.Put([]byte("b_" + b.hashCurr), bs, nil); err != nil {
		log.Fatal("block can not be saved")
	}
	// 将最后的区块哈希设置为当前区块
	bc.lastHash = b.hashCurr
	// 将最后的区块哈希存储到数据库中
	if err := bc.db.Put([]byte("lastHash"), []byte(b.hashCurr), nil); err != nil {
		log.Fatal("lastHas can not be saved")
	}
	return bc
}

// 通过hash获取区块
func (bc *BlockChain) GetBlock(hash  Hash) (*Block, error) {
	// 从数据库中读取对应的区块
	data, err := bc.db.Get([]byte("b_" + hash), nil)
	if err != nil {
		return nil, err
	}
	// 反序列化
	b, err := BlockUnSerialize(data)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// 迭代展示区块的方法
func (bc *BlockChain) Iterate() {
	// 最后的哈希
	for hash := bc.lastHash; hash != ""; {
		// 得到区块
		b, err := bc.GetBlock(hash)
		if err != nil {
			log.Fatalf("Block <%s> is not exists.", hash)
		}
		fmt.Println("HashCurr:", b.hashCurr)
		fmt.Println("TXs:", b.txs)
		fmt.Println("Time:", b.header.time.Format(time.UnixDate))
		fmt.Println("HashPrev:", b.header.hashPrevBlock)
		fmt.Println()

		hash = b.header.hashPrevBlock
	}
}


