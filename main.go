package main

import (
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
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.BlockchainAddress())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	fmt.Printf("Singature :     %s\n", t.GenerateSignature().String())
}
