package main

import ("fmt"
		"strconv"
		"bytes"
		"crypto/sha256"
		"time"
	   )
// Block : 시간 거래기록 이전 해쉬 그리고 지금 해쉬를 각각 정수 또는 슬라이스 바이트로 기록
type Block struct{
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

// SetHash: 앞에 먼저 변수를 정의하면 거기에 대한 메서드를 지정할 수 있음 // Hash 메서드를 지정하기 위함
// 먼저 timestmap 를 정수형에서 바이트 슬라이스로 변환
// 변환된 timestmap와 데이터 그리고 이전해쉬를 {} 아무 붙임없이 그냥 이음 바이트 슬라이스 형태로
// sha256 을 통해 해쉬화
// hash 를 b.Hash 메서드로 지정
func (b *Block) SetHash() {
	timestmap := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestmap}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
// 새로운 블럭은 데이터 그리고 이전 블럭 해쉬로 만들어짐
// 시간은 지금 시간을 가져와 기록하고 나머지 데이터를 채우고 hash 는 공란으로 남겨둠
// 블럭을 리턴
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{}}
	// 여기서 새로운 저장공간이 할당됨(새로운 &Block 이 만들어지면서 저장공간 할당) block 은 포인터가됨!!!!!!!!!!
	block.SetHash()
	return block
}
// Block 포인터 데이터의 집합 
type BlockChain struct{
	blocks []*Block
}

func (bc *BlockChain) AddBlock(data string){
	prevBlock := bc.blocks[len(bc.blocks) -1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block{
	return NewBlock("Gensis Block", []byte{})
	// NewGenesisBlock의 저장공간은 NewBlock 에서 할당됨
}

func NewBlockchain() *BlockChain{
	return &BlockChain{[]*Block{NewGenesisBlock()}}
	// 여기서 새로운 블록체인의 저장공간이 할당됨 !!!!!!!!!!!!!!!!!!
}

func main() {
	bc := NewBlockchain()
	
	bc.AddBlock("김재영 뒤져")
	bc.AddBlock("안뒤질랭")
	fmt.Println(*bc.blocks[2])
	bc.blocks[0].Timestamp = 1234567890
	// 작업증명이 없기 때문에 쉽게 변형이 가능함
	for _,block := range bc.blocks{
		fmt.Printf("Prev. hash: %d\n",block.PrevBlockHash)
	}
}