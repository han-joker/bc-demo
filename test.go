package main

import (
	"fmt"
	"github.com/han-joker/bc-demo/blockchain"
)

func main() {
	b := blockchain.NewBlock("", "Gensis Block.")
	fmt.Println(b)
}
