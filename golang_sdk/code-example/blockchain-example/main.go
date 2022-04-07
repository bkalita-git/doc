package main

import (
	"blockchain-example/blockchain"
	"encoding/binary"
	"fmt"
	"strconv"
	"unsafe"
)

type wonder struct {
	x int
}

func main() {
	var x = 1
	fmt.Printf("%d---%x", unsafe.Sizeof(x), binary.BigEndian)
	return
	var chain = blockchain.InitBlockChain()
	chain.AddBlock("this is 1")
	chain.AddBlock("this is 2")
	chain.AddBlock("this is 3")

	for _, block := range chain.Blocks {
		fmt.Printf("\nprevious Hash %x", block.PrevHash)
		fmt.Printf("\nData in blocks %s", block.Data)
		fmt.Printf("\nHash %x", block.Hash)

		var pow = blockchain.NewProof(block)
		fmt.Printf("\nPoW %s", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
