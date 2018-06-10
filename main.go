package main

import(
	"crypto/sha256"
	"encoding/hex"
	"time"
	"strings"
)

type Data struct {
	Message string
}

type Block struct {
	Index     int
	Timestamp string
	Data 	  Data
	Hash      string
	PrevHash  string
	Nonce 	  string
}

var Blockchain []Block
var TRANSACTION_PER_BLOCK = 20

func main() {
	data := Data{Message: "salut"}

	myBlock := Block {
		Index: 0,
		Timestamp: "test",
		Data: data,
		Hash: "123123123",
		PrevHash: "02392392",
	}

	proofOfWork(myBlock)

	println(myBlock.Nonce)

	println(myBlock.PrevHash)
}

func calculateHash(block Block) string {
	record := string(block.Index) + string(block.Timestamp) + string(block.Data.Message) + string(block.PrevHash)

	h := sha256.New()

	h.Write([]byte(record))

	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data Data) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func proofOfWork(newBlock Block) Block {
	complete := false
	var n uint64
	n = 0
	h := sha256.New()
	for complete == false {

		stringToBeHashed := newBlock.Data.Message + string(n)

		h.Write([]byte(stringToBeHashed))

		hashed := h.Sum(nil)

		hashString := hex.EncodeToString(hashed)

		println(hashString)

		if strings.HasPrefix(hashString, "000000") {
			println("Proof of work completed")
			println("hash is %s", hashString)

			complete = true
			newBlock.Nonce = string(hashed)
		}

		n += 1
		//println(n)
	}

	return newBlock
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

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}