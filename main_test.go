package main

import (
	"fmt"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	block := createBlock(1)
	result := calculateHash(block)
	if result == "" {
		t.Errorf("Expecting hash, but got %s", result)
	}
}

func TestGenerateBlock(t *testing.T) {
	prevBlockIndex := 2
	bpm := 60
	block := createBlock(prevBlockIndex)
	nextBlock, err := generateBlock(block, bpm)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}

	if nextBlock.BPM != bpm {
		t.Errorf(fmt.Sprintf("Expected %d BPM, but got %d", bpm, nextBlock.BPM))
	}

}

func TestIsBlockValid(t *testing.T) {
	block := createBlock(1)
	nextBlock, _ := generateBlock(block, 60)

	valid := isBlockValid(nextBlock, block)
	if valid != true {
		t.Errorf("Expecting true, but got %t", valid)
	}
}

func createBlock(index int) Block {
	return Block{
		Index:     index,
		Timestamp: "12:00pm",
		BPM:       50,
		Hash:      "67890",
		PrevHash:  "12345",
	}
}
