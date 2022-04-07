package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// func (b *Block) DeriveHash() {
// 	var info = bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	var hash = sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

func CreateBlock(data string, prevHash []byte) *Block {
	var block = &Block{Hash: []byte{}, Data: []byte(data), PrevHash: prevHash, Nonce: 0}

	var pow = NewProof(block)
	var nonce, hash = pow.Run()
	// block.DeriveHash()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	var prevBlock = chain.Blocks[len(chain.Blocks)-1]
	var new = CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
