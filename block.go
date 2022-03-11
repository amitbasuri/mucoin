package mucoin

import (
	"bytes"
	"crypto/sha256"
)

const genesis = "Genesis"

type block struct {
	hash     []byte
	data     []byte
	prevHash []byte
}

func (b *block) generateHash() {
	info := bytes.Join([][]byte{b.data, b.prevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *block {
	block := &block{[]byte{}, []byte(data), prevHash}
	block.generateHash()
	return block
}

func GenesisBlock() *block {
	return CreateBlock(genesis, []byte{})
}
