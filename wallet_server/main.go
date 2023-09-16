package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("WALLET SERVER :")
}

func main() {

	port := flag.Uint("port", 8080, "TCP Port Number for wallet Server")
	gateway := flag.String("gateway", "https://127.0.0.1:3000", "BlockChain Gateway")
	flag.Parse()

	app := NewWalletServer(uint16(*port), *gateway)
	app.Run()
}
