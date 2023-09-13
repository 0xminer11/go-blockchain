package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactios  []*Transaction
}

type Transaction struct {
	senderBlockchainAddress  string
	reciverBlockchainAddress string
	value                    float32
}

func NewTransaction(_sender string, _reciver string, _value float32) *Transaction {
	return &Transaction{_sender, _reciver, _value}
}

func (t *Transaction) print() {
	fmt.Printf("SenderAddress        %s\n", t.senderBlockchainAddress)
	fmt.Printf("ReciverAddress       %s\n", t.reciverBlockchainAddress)
	fmt.Printf("Value        		%f\n", t.value)
}

func (t *Transaction) MarshelJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress  string  `json:"sender_blockchain_address"`
		ReciverBlockchainAddress string  `json:"reciver_blockchain_address"`
		Value                    float32 `json:"value"`
	}{
		SenderBlockchainAddress:  t.senderBlockchainAddress,
		ReciverBlockchainAddress: t.reciverBlockchainAddress,
		Value:                    t.value,
	})
}
func NewBlock(nonce int, previousHash [32]byte, txs []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactios = txs
	return b
}

func (b *Block) print() {
	fmt.Printf("Nonce         %d\n", b.nonce)
	fmt.Printf("timestamp     %d\n", b.timestamp)
	fmt.Printf("previousHash  %x\n", b.previousHash)
	//fmt.Printf("transactions  %s\n", b.transactios)
	for _, t := range b.transactios {
		t.print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func newBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.createBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) print() {
	for i, block := range bc.chain {
		fmt.Println("Chain    ", i)
		block.print()
	}
}

func (bc *Blockchain) createBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) AddTransaction(_sender string, _reciver string, _value float32) {
	t := NewTransaction(_sender, _reciver, _value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	sum := sha256.Sum256([]byte(m))
	fmt.Printf("%x", sum)
	return sum

	//fmt.Println("Hash", sha256.Sum256([]byte("hello")))
	//return sha256.Sum256([]byte("hello"))
}

func init() {
	log.SetPrefix("Blockchain :")
}

func main() {
	bc := newBlockchain()
	log.Println("Block Mined..✅")
	bc.print()

	hash := bc.LastBlock().Hash()
	bc.createBlock(1, hash)
	bc.print()

	bc.AddTransaction("A", "B", 1.0)
	bc.AddTransaction("A", "B", 2.0)
	hash = bc.LastBlock().Hash()
	bc.createBlock(2, hash)
	bc.print()

	hash = bc.LastBlock().Hash()
	bc.createBlock(3, hash)
	bc.print()

}
