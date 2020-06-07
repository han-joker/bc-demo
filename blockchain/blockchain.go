package blockchain

import (
	"fmt"
	"time"
)

type BlockChain struct {
	lastHash Hash // 最后一个区块的哈希
	blocks map[Hash]*Block // 全部区块信息，由区块哈希作为 key 来检索
}

// 构造方法
func NewBlockchain() *BlockChain {
	bc := &BlockChain{
		blocks:   map[Hash]*Block{},
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
	bc.blocks[b.hashCurr] = b
	// 将最后的区块哈希设置为当前区块
	bc.lastHash = b.hashCurr

	return bc
}

// 迭代展示区块的方法
func (bc *BlockChain) Iterate() {
	// 最后的哈希
	for hash := bc.lastHash; hash != ""; {
		b := bc.blocks[hash]
		fmt.Println("HashCurr:", b.hashCurr)
		fmt.Println("TXs:", b.txs)
		fmt.Println("Time:", b.header.time.Format(time.UnixDate))
		fmt.Println("HashPrev:", b.header.hashPrevBlock)
		fmt.Println()

		hash = b.header.hashPrevBlock
	}
}


