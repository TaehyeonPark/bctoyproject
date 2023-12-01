package main

import (
	"fmt"
	"os"

	"helloBlockchain/crypto"
	"helloBlockchain/storage"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("[Error] Cannot load .env")
	}

	config := &storage.Config{
		Host: os.Getenv("DATABASE_HOST"),
		Port: os.Getenv("DATABASE_PORT"),
		Password: os.Getenv("DATABASE_PW"),
		User: os.Getenv("DATABASE_USER"),
		DBName: os.Getenv("DATABASE_NAME"),
		SSLMode: os.Getenv("DATABASE_SSLMODE")}

	db, err := storage.NewConnection(config)
	
	if err != nil {
		fmt.Println("[Error] Cannot connect to DB")
	}

}

func demo(db *gorm.DB) {
	bc := crypto.NewBlockchain(db)
	bc.AddBlock("Send [1] Ara from {me} to {Taehyeon}")
	bc.AddBlock("Send [2] Ara from {me} to {Taehyeon}")
	
	for _, block := range bc.GetBlocks() {
		fmt.Printf("Timestamp	: %d\n", block.Timestamp)
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