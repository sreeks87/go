package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/sreeks87/go/blockchain/blockchain"
	block "github.com/sreeks87/go/blockchain/blockchain"
)

type CommandLine struct {
	blockchain *block.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage :")
	fmt.Println("add -block BLOCK_DATA -> adds a new blockto the chain")
	fmt.Println("print - prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		// this is used for badger to GC and shutdown properly.
		runtime.Goexit()
	}
}
func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()
	for {
		block := iter.Next()

		fmt.Println("Previous hash :- ", block.PrevHash)
		fmt.Println("Block data :- ", block.Data)
		fmt.Println("Block hash :- ", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("POW %s", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block Data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()

	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}
	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Block added.")
}

func main() {
	defer os.Exit(0)
	chain := block.InitBlockChain()
	defer chain.Database.Close()
	cli := CommandLine{chain}
	cli.Run()

}
