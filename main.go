package main

import (
	"blockchain.com/block"
	"blockchain.com/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain :")
}

func main() {
	log.Println("Hi")

	w := wallet.NewWallet()
	w1 := wallet.NewWallet()
	w2 := wallet.NewWallet()

	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.BlockchainAddress())

	t := wallet.NewTransaction(w1.PrivateKey(), w1.PublicKey(), w1.BlockchainAddress(), w2.BlockchainAddress(), 1.0)
	//fmt.Printf("Singature :     %s\n", t.GenerateSignature().String())

	bc := block.NewBlockchain(w.BlockchainAddress())
	isadded := bc.AddTransaction(w1.BlockchainAddress(), w2.BlockchainAddress(), 1.0, w1.PublicKey(), t.GenerateSignature())
	fmt.Println("IsAdded", isadded)

	bc.Mining()

	print(bc)
}
