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
	fmt.Println(w.PublicKeyStr())
}
