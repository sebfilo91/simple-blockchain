package main

import (
	"github.com/simple-blockchain/node"

	"github.com/simple-blockchain/core"
)


var Blockchain []core.Block
var nodes []node.Node

//var connections []*Connection

func main() {
	otherNode := node.Node{}
	otherNode.Address = "localhost"

	nodes = append(nodes, otherNode)

	/*data := Data{Message: "salut"}

	myBlock := Block{
		Index:     0,
		Timestamp: time.Now(),
		Data:      data,
		Hash:      "123123123",
		PrevHash:  "02392392",
	}*/

	//proofOfWork(myBlock)
}

func connectToNodes() {
	/*conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}*/
}



func broadcastNewblock() {

}



func replaceChain(newBlocks []core.Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

