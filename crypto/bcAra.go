/* 테스트를 위해 접근제한을 외부 사용 가능으로 둠 */
package crypto

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

const targetBits = 1			// 채굴 난이도
const maxNonce = math.MaxInt64	// nonce overflow 방지
const nonceStartAtZero = 0

type Block struct {
	Timestamp		int64
	Data			[]byte
	PrevBlockHash	[]byte
	Hash			[]byte
	Nonce			int
}

func (b *Block) Serialize() []byte {
	var buffer bytes.Buffer
	
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Println("Error in Serialize : ", err)
	}

	return buffer.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var b *Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&b)
	if err != nil {
		fmt.Println("Error in DeserializeBlock : ", err)
	}

	return b
}

/* Genesis Block initialization */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

/* Block constructor */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data: []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash: []byte{},
		Nonce: nonceStartAtZero,
	}
	// [Deprecated] --> block.SetHash()
	// 이제는 채굴이 필요함
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

/* Blockchain constructor */
func NewBlockChain() *Blockchain {
	return &Blockchain{blocks: []*Block{NewGenesisBlock()}}
}

type ProofOfWork struct {
	block *Block
	target *big.Int	// 해시값과 비교될 조건
}

/* nonce : hashcache에서 counter와 동일한 역할 */
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := nonceStartAtZero
	
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	
	for nonce < maxNonce {	// nonce의 overflow를 검사함
		data := pow.PrepareData(nonce)		// 1. 데이터 생성
		hash = sha256.Sum256(data)			// 2. 해시
		fmt.Printf("\r%x", hash)
		
		hashInt.SetBytes(hash[:])			// 3. bigInt로 변환
		if hashInt.Cmp(pow.target) < 0 {	// 4. target(조건)과 비교
			break
			} else {
				nonce++
			}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) < 0
}

/* ProofOfWork constructor */
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	/* targetBits(난이도)가 높을수록 target(조건)은 작아짐(어려워짐) */
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{block: b, target: target}
	return pow
}


func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

type Repository struct {
}

/* [DEPRECATED] */
// func (b *Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)	// Hash Algorithm : SHA-256
// 	b.Hash = hash[:]
// }