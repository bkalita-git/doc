package blockchain

// proof of work
//powcomutation >validation time

//Requirements
//the first few bytes must contain 0s

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	var target = big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	var pow = &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	var data = bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)), //casting
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	var buff = new(bytes.Buffer)
	var err = binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	var nonce = 0
	for nonce < math.MaxInt64 {
		var data = pow.InitData(nonce)
		var hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	var data = pow.InitData(pow.Block.Nonce)
	var hash = sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1

}
