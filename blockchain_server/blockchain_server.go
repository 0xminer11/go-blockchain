package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type BlockchaonServer struct {
	port uint16
}

func NewBlockchainServer(port uint16) *BlockchaonServer {
	return &BlockchaonServer{port: port}
}

func (bsc *BlockchaonServer) Port() uint16 {
	return bsc.port
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Blockchain")
}

func (bsc *BlockchaonServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bsc.Port())), nil))
}
