package node

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"github.com/simple-blockchain/core"
)

type Node struct {
	Address string
}

func proofOfWork(newBlock core.Block) core.Block {
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
			// broadcastNewblock
		}

		n += 1
	}

	return newBlock
}
