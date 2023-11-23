package main

import (
	"fmt"
	"helloBlockchain/crypto"
)

func main() {
	bc := crypto.NewBlockChain()

	bc.AddBlock("Send [1] Ara from {me} to {Taehyeon}")
	bc.AddBlock("Send [2] Ara from {me} to {Taehyeon}")
	for _, block := range bc.GetBlocks() {
		fmt.Printf("Prev Hash	: %x\n", block.PrevBlockHash)
		fmt.Printf("Data		: %s\n", block.Data)
		fmt.Printf("Hash		: %x\n", block.Hash)
		fmt.Printf("Nonce		: %d\n", block.Nonce)
		pow := crypto.NewProofOfWork(block)
		pow.PrepareData(block.Nonce)
		isValid := pow.Validate()
		fmt.Printf("IsValid		: %t\n", isValid)
		fmt.Println()
	}


}