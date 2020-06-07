package main

import (
	"flag"
	"fmt"
	"github.com/han-joker/bc-demo/blockchain"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"strings"
)

// 命令行工具
func main() {
	// # 初始化数据库
	// 数据库连接
	dbpath := "data"
	db, err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 释放数据库连接
	defer db.Close()

	// 初始化区块链
	bc := blockchain.NewBlockchain(db)
	// 添加创世区块
	bc.AddGensisBlock()

	// 初始化第一个命令参数
	arg1 := ""
	// 若用户指定了参数，则第一个用户参数为命令参数
	if len(os.Args) >= 2 {
		arg1 = os.Args[1]
	}
	// 基于命令参数，执行对应的功能
	switch strings.ToLower(arg1) {
	case "create:block":
		// 为 createblock 命令增加一个 flag 集合。标志集合
		// 错误处理为，一旦解析失败，则 exit
		fs := flag.NewFlagSet("create:block", flag.ExitOnError)
		// 在集合中，添加需要解析的 flag 标志
		txs := fs.String("txs", "", "")
		// 解析命令行参数,
		fs.Parse(os.Args[2:])
		// 完成区块的创建
		bc.AddBlock(*txs)
	case "show":
		bc.Iterate()
	case "init":
		// 清空
		bc.Clear()
		// 添加创世区块
		bc.AddGensisBlock()
	case "help":
		fallthrough
	default:
		Usage()
	}
}

// 输出 bcli 的帮助信息
func Usage() {
	fmt.Println("bcli is a tool for Blockchain.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Printf("\t%s\t%s\n", "bcli createblock -txs=<txs>", "create block on blockchain")
	fmt.Printf("\t%s\t\t\t%s\n", "bcli init", "initial blockchain")
	fmt.Printf("\t%s\t\t\t%s\n", "bcli help", "help info for bcli")
	fmt.Printf("\t%s\t\t\t%s\n", "bcli show", "show blocks in chain.")
}