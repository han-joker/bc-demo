package main

import (
	"github.com/han-joker/bc-demo/blockchain"
)

func main() {
	//b := blockchain.NewBlock("", "Gensis Block.")
	//fmt.Println(b)

	// 区块链测试
	bc := blockchain.NewBlockchain()
	bc.AddGensisBlock()
	bc.
		AddBlock("First Block").
		AddBlock("Second Block")

	bc.Iterate()
}
