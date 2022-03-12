package mucoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/big"
)

const Difficulty = 6

type ProofOfWork struct {
	target *big.Int
}

func NewProofOfWork() *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{
		target: target,
	}

	return pow
}

func (pow *ProofOfWork) InitData(block *block, nonce int) []byte {
	data := bytes.Join([][]byte{block.data, block.prevHash, ToHex(int64(nonce))}, []byte{})
	return data
}

func (pow *ProofOfWork) Run(block *block) (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := pow.InitData(block, nonce)
		hash := sha256.Sum256(data)
		//fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate(block *block) bool {
	var intHash big.Int

	data := pow.InitData(block, block.nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	if err := binary.Write(buff, binary.BigEndian, num); err != nil {
		panic(err)
	}
	return buff.Bytes()
}
