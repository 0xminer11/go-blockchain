package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain")
}

func main() {
	log.Println("Starting...")
	port := flag.Uint("port", 3000, "TCP Port for Blockchain Server")
	flag.Parse()
	fmt.Println(*port)
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
