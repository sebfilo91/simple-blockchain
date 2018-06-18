package core

import "time"

type Block struct {
	Index     int
	Timestamp time.Time
	Data      Data
	Hash      string
	PrevHash  string
	Nonce     string
}
type Data struct {
	Message string
}

var TRANSACTION_PER_BLOCK = 20


func generateBlock(oldBlock Block, data Data) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}