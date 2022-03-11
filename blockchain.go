package mucoin

type BlockChain struct {
	blocks []*block
}

func (chain *BlockChain) AddBlock(data string) {
	block := CreateBlock(data, chain.blocks[len(chain.blocks)-1].hash)
	chain.blocks = append(chain.blocks, block)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*block{GenesisBlock()},
	}
}
