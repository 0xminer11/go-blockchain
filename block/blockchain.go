package block

import (
	"blockchain.com/utils"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "svs"
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

func NewBlockchain(blockchainAddress string, port uint16) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.createBlock(0, b.Hash())
	bc.port = port
	return bc
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		previousHash string         `json:"previousHash"`
		Transaction  []*Transaction `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		previousHash: fmt.Sprintf("%x", b.previousHash),
		Transaction:  b.transactios,
	})
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
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
	port              uint16
}

func newBlockchain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.createBlock(0, b.Hash())
	return bc
}
func (bc *Blockchain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []*Block `json:"chain"`
	}{
		Blocks: bc.chain,
	})
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

func (bc *Blockchain) AddTransaction(_sender string, _receiver string, _value float32, senderPublicKey *ecdsa.PublicKey, s *utils.Signature) bool {

	t := NewTransaction(_sender, _receiver, _value)

	if _sender == MINING_SENDER {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	}

	if bc.verifyTransactionSignature(senderPublicKey, s, t) {
		bc.transactionPool = append(bc.transactionPool, t)
		return true
	} else {
		log.Println("TX : failed to verify")
	}
	return false
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

func (bc *Blockchain) CopyTransaction() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(
				t.senderBlockchainAddress,
				t.reciverBlockchainAddress,
				t.value))
	}
	return transactions
}

//func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, txs []*Transaction, difficulty int) bool {
//	zeros := strings.Repeat("0", difficulty)
//	guessBlock := Block{nonce, previousHash, 0, txs}
//	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
//	return guessHashStr[:difficulty] == zeros
//}

func (bc *Blockchain) ProofOfStake() int {
	//transactions := bc.CopyTransaction()
	previousnonce := bc.LastBlock().nonce
	return previousnonce + 1

	//for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
	//	nonce += 1
	//}
	//return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, 1.0, nil, nil)
	nonce := bc.ProofOfStake()
	prevHash := bc.LastBlock().Hash()
	bc.createBlock(nonce, prevHash)
	log.Println(" Mining Successful Block Created ⛏️⛏️⛏️")
	return true
}

func (bc *Blockchain) verifyTransactionSignature(sender *ecdsa.PublicKey, s *utils.Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	return ecdsa.Verify(sender, h[:], s.R, s.S)
}

func init() {
	log.SetPrefix("Blockchain :")
}

func main() {
	//bc := newBlockchain()
	//log.Println("Block Mined..✅")
	//bc.print()
	////nonce := bc.ProofOfWork()
	//hash := bc.LastBlock().Hash()
	//bc.createBlock(1, hash)
	//bc.print()
	//nonce := bc.ProofOfStake()
	//bc.AddTransaction("A", "B", 1.0)
	//bc.AddTransaction("A", "B", 2.0)
	//hash = bc.LastBlock().Hash()
	//bc.createBlock(nonce, hash)
	//bc.print()
	//nonce = bc.ProofOfStake()
	//hash = bc.LastBlock().Hash()
	//bc.createBlock(nonce, hash)
	//bc.print()

}
