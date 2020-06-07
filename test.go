package main

import (
	"github.com/han-joker/bc-demo/blockchain"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	// 数据库连接
	dbpath := "data"
	db, err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 释放数据库连接
	defer db.Close()

	// 区块链测试
	bc := blockchain.NewBlockchain(db)
	// 添加创世区块
	bc.AddGensisBlock()
	// 添加新区块
	bc.
		AddBlock("First Block").
		AddBlock("Second Block")

	bc.Iterate()
}
