package main

import (
	"blockchain.com/block"
	"blockchain.com/wallet"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type BlockchaonServer struct {
	port uint16
}

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

func NewBlockchainServer(port uint16) *BlockchaonServer {
	return &BlockchaonServer{port: port}
}

func (bsc *BlockchaonServer) Port() uint16 {
	return bsc.port
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Blockchain")
}

func (bsc *BlockchaonServer) GetChain(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bsc.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))

	default:
		log.Println("ERROR :Invalid HTTP Method")
	}
}

func (bsc *BlockchaonServer) Run() {
	http.HandleFunc("/", bsc.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bsc.Port())), nil))
}

func (bsc *BlockchaonServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]

	if !ok {
		minerWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minerWallet.BlockchainAddress(), bsc.Port())
		cache["blockchain"] = bc
		fmt.Printf("PrivateKey    %v", minerWallet.PrivateKeyStr())
		fmt.Printf("PublicKey    %v", minerWallet.PublicKeyStr())
		fmt.Printf("BlockchainAddress    %v", minerWallet.BlockchainAddress())
	}
	return bc
}
