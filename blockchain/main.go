package main

import (
	"fmt"
	"strconv"

	block "github.com/sreeks87/go/blockchain/blockchain"
)

func main() {
	chain := block.InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, b := range chain.Blocks {
		fmt.Println("Previous hash :- ", b.PrevHash)
		fmt.Println("Block data :- ", b.Data)
		fmt.Println("Block hash :- ", b.Hash)

		pow := block.NewProof(b)
		fmt.Printf("POW %s", strconv.FormatBool(pow.Validate()))
		fmt.Println()

	}
}
