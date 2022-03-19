package chain

import (
	"github.com/asynched/blockchain/domain/block"
)

type Blockchain struct {
	Blocks []block.Block `json:"blocks"`
}

func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, previousBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

}

func NewBlockChain() Blockchain {
	return Blockchain{
		Blocks: []block.Block{block.NewGenesisBlock()},
	}
}
