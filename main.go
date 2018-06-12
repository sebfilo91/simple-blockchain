package main

import(
	"crypto/sha256"
	"encoding/hex"
	"time"
	"strings"
	"net"
	"github.com/tyler-smith/go-bip39"
 	"github.com/tyler-smith/go-bip32"
)


type Data struct {
	Message string
}

type Block struct {
	Index     int
	Timestamp time.Time
	Data 	  Data
	Hash      string
	PrevHash  string
	Nonce 	  string
}

type Node struct {
	Address string
}

var Blockchain []Block
var TRANSACTION_PER_BLOCK = 20
var nodes []Node
var connections []*Connection

func main() {
	otherNode := Node{}
	otherNode.Address = "localhost"

	nodes = append(nodes, otherNode)

	data := Data{Message: "salut"}

	myBlock := Block {
		Index: 0,
		Timestamp: time.Now(),
		Data: data,
		Hash: "123123123",
		PrevHash: "02392392",
	}

	proofOfWork(myBlock)

	println(myBlock.Nonce)

	println(myBlock.PrevHash)
}

func connectToNodes() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp.String() + string(block.Data.Message) + string(block.PrevHash)

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

		if strings.HasPrefix(hashString, "000") {
			println("Proof of work completed")
			println("hash is %s", hashString)

			complete = true
			newBlock.Nonce = string(hashed)
			broadcastNewblock()
		}

		n += 1
		//println(n)
	}

	return newBlock
}

func broadcastNewblock() {

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

func createMnemonic() {
	  // Generate a mnemonic for memorization or user-friendly seeds
	  entropy, _ := bip39.NewEntropy(256)
	  mnemonic, _ := bip39.NewMnemonic(entropy)
	
	  // Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	  seed := bip39.NewSeed(mnemonic, "Secret Passphrase")
	
	  masterKey, _ := bip32.NewMasterKey(seed)
	  publicKey := masterKey.PublicKey()
	
	  // Display mnemonic and keys
	  fmt.Println("Mnemonic: ", mnemonic)
	  fmt.Println("Master private key: ", masterKey)
	  fmt.Println("Master public key: ", publicKey)
}

func validateTransaction() {
	// Check key from sender
}