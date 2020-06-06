package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Hash = string

const nodeVersion = 0

// 区块主体
type Block struct {
	header    BlockHeader
	txs       string // 交易列表
	txCounter int    // 交易计数器
	hashCurr  Hash   // 当前区块哈市值缓存
}

// 区块头
type BlockHeader struct {
	version        int
	hashPrevBlock  Hash      // 前一个区块的 Hash
	hashMerkleRoot Hash      // 默克尔树的哈希节点
	time           time.Time // 区块的创建时间
	bits           int       // 难度相关
	nonce          int       // 挖矿计数
}

// 构造区块
func NewBlock(prevHash Hash, txs string) *Block {
	// 实例化Block
	b := &Block{
		header: BlockHeader{
			version:       nodeVersion,
			hashPrevBlock: prevHash, // 设置前面的区块哈希
			time:          time.Now(),
		},
		txs:       txs, // 设置数据
		txCounter: 1,   // 计数交易
	}
	// 计算设置当前区块的哈希
	b.SetHashCurr()
	return b
}

// 头信息的字符串化
func (bh *BlockHeader) Stringify() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.hashPrevBlock,
		bh.hashMerkleRoot,
		bh.time.UnixNano(), // 得到时间戳，nano 级别
		bh.bits,
		bh.nonce,
	)
}

// 设置当前区块 hash
func (b *Block) SetHashCurr() *Block {
	// 生成头信息的拼接字符串
	headerStr := b.header.Stringify()
	// 计算 hash 值
	b.hashCurr = fmt.Sprintf("%x", sha256.Sum256([]byte(headerStr)))

	return b
}
