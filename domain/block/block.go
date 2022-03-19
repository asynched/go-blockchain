package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp         int64  `json:"timestamp"`
	Data              []byte `json:"data"`
	PreviousBlockHash []byte `json:"previous_block_hash"`
	Hash              []byte `json:"hash"`
}

func makeHash(b *Block) []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreviousBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	return hash[:]
}

func NewBlock(data string, previousBlockHash []byte) Block {
	block := Block{
		Timestamp:         time.Now().Unix(),
		Data:              []byte(data),
		PreviousBlockHash: previousBlockHash,
		Hash:              []byte{},
	}

	block.Hash = makeHash(&block)

	return block
}

func NewGenesisBlock() Block {
	return NewBlock("Genesis Block", []byte{})
}
