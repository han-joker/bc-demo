package blockchain

import (
	"errors"
	"github.com/han-joker/bc-demo/block"
)

type BCIterator struct {
	// 关联需要迭代的对象
	bc *BlockChain
	// 当前迭代的区块hash
	currHash block.Hash
}

// 构造迭代器
func NewBCIterator(bc *BlockChain) *BCIterator {
	return &BCIterator{
		bc:       bc,
		currHash: bc.lastHash,
	}
}

// for b, err := bci.Next(); err == nil {
// 		循环主体
// }
func (bci *BCIterator) Next() (*block.Block, error) {
	// 如果currHash == "" 表示也没有区块了，不用再操作数据库
	if bci.currHash == "" {
		return nil, errors.New("")
	}

	// 当前hash在数据库中，可以获取内容，表示没有错误，可以循环
	block, err := bci.bc.GetBlock(bci.currHash)
	if err != nil {
		return nil, err
	}
	// 得到前一个区块的 hash
	bci.currHash = block.GetHashPrevBlock()
	// 返回当前区块和没有错误
	return block, nil
}


//// 用于判断是否可以继续迭代 for bci.Valid()
//func (bci *BCIterator) Valid() bool {
//	return  true
//}

// for bci.Valid() {
// 		b := bci.Next()
// }


