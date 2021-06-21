package BC

// import(
// 	"bytes"
// 	"crypto/sha256"

// )

type Block struct{
	Hash []byte
	Data []byte
	PrevHash []byte
}


type BlockChain struct{
	Blocks []Block
}

// func (b *Block) DeriveHash(){
// 	info := bytes.Join([][]byte{b.Data,b.PrevHash},[]byte{})

// 	hash := sha256.Sum256(info)

// 	b.Hash = hash[:]
// }

// func CreateBlock(data []byte, prevHash []byte) *Block {
// 	block := &Block{[]byte{}, data, prevHash}
	
// 	block.DeriveHash()
// 	return block
// }

// func (chain *BlockChain) AddBlock(data []byte){
// 	prevBlock := chain.Blocks[len(chain.Blocks)-1]
// 	new := CreateBlock(data, prevBlock.Hash)
// 	chain.Blocks = append(chain.Blocks, new)
// }

// func Genesis() *Block{
// 	str := "Genesis"
// 	return CreateBlock([]byte(str),[]byte{})
// }


// func InitBlockChain() *BlockChain {
// 	return &BlockChain{[]*Block{Genesis()}}
// }

