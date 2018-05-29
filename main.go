package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	/*
		"encoding/json"
		"fmt"
		"io"
		"log"
		"net/http"
		"os"
		"time"

		"github.com/davecgh/go-spew/spew"
		"github.com/gorilla/mux"
		"github.com/joho/godotenv"
	*/)

type Block struct {
	// Position of the data record in the blockchain
	Index int

	// Automatically determined and is the time the data is written
	Timestamp string

	// Pulse rate
	BPM int

	// SHA256 identifier representing  this data record
	Hash string

	// SHA256 identifie of the previous record in the chain
	PrevHash string
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp +
		string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func main() {
	fmt.Println("vim-go")
}
