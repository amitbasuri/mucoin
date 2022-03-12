package mucoin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitBlockChain(t *testing.T) {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	assert.Equal(t, chain.blocks[3].prevHash, chain.blocks[2].hash, "prev block hash must be equal to next block prevHash")
	assert.Equal(t, chain.blocks[2].prevHash, chain.blocks[1].hash, "prev block hash must be equal to next block prevHash")
	assert.Equal(t, chain.blocks[1].prevHash, chain.blocks[0].hash, "prev block hash must be equal to next block prevHash")

	pow := NewProofOfWork()
	assert.Equal(t, pow.Validate(chain.blocks[3]), true, "block must be valid")
	assert.Equal(t, pow.Validate(chain.blocks[2]), true, "block must be valid")
	assert.Equal(t, pow.Validate(chain.blocks[1]), true, "block must be valid")
	assert.Equal(t, pow.Validate(chain.blocks[0]), true, "block must be valid")

}
