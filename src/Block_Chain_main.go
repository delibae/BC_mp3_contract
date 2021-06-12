package main

import ("fmt"
		"strconv"
		"bytes"
		"crypto/sha256"
		"time"
	   )

type Block struct{
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

func (b *Block) SetHash() {
	timestmap := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestmap}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{}}
	block.SetHash()
	return block
}

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
}

func NewBlockchain() *BlockChain{
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockchain()
	
	bc.AddBlock("김재영 뒤져")
	bc.AddBlock("안뒤질랭")
	fmt.Println(*bc.blocks[2])
	for _,block := range bc.blocks{
		fmt.Printf("Prev. hash: %d\n",block.PrevBlockHash)
	}
}