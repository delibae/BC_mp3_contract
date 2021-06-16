package main

import("math/big"
	  "fmt"
	  
	  )

const targetBits = 24

type Block struct{
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
	Nonce int 
}



type POW struct{
	block *Block
	target *big.Int
}

func NewPOW(b *Block) *POW{
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &POW{b,target}
	return pow
}

func main(){
	fmt.Println("Hellow world")
}
