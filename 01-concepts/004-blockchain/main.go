package main

import (
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	blockChain := NewBlockChain()
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	nonce := blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("E", "F", 3.0)
	previousHash = blockChain.LastBlock().Hash()
	nonce = blockChain.ProofOfWork()
	blockChain.CreateBlock(nonce, previousHash)
	blockChain.Print()

}
