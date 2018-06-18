package core

import (
	"crypto/sha256"
	"encoding/hex"
)

type Blockchain struct {

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

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp.String() + string(block.Data.Message) + string(block.PrevHash)

	h := sha256.New()

	h.Write([]byte(record))

	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}