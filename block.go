package mucoin

const genesis = "Genesis"

type block struct {
	hash     []byte
	data     []byte
	prevHash []byte
	nonce    int
}

func CreateBlock(data string, prevHash []byte) *block {
	block := &block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProofOfWork()
	block.nonce, block.hash = pow.Run(block)
	return block
}

func GenesisBlock() *block {
	return CreateBlock(genesis, []byte{})
}
