package main

import (
	"time"
)

type Block struct {
	number   uint32
	prevHash string
	uncles   []*Block
	coinbase string

	difficulty   uint32
	time         int64
	nonce        uint32
	transactions []*Transaction
	extra        string
}

// NewBlock returns Block
func NewBlock(raw []byte) *Block {
	block := &Block{}
	block.UnmarshalRlp(raw)
	return block
}

func CreateBlock(
	transactions []*Transaction,
) *Block {
	block := &Block{
		transactions: transactions,
		number:       1,
		prevHash:     "1234",
		coinbase:     "me",
		difficulty:   10,
		nonce:        0,
		time:         time.Now().Unix(),
	}

	return block
}

// Hash returns string
func (block *Block) Hash() string {
	return Sha256Hex(block.MarshalRlp())
}

// MarshalRlp returns []byte
func (block *Block) MarshalRlp() []byte {
	encTx := make([]string, len(block.transactions))
	for i, tx := range block.transactions {
		encTx[i] = string(tx.MarshalRlp())
	}

	header := []interface{}{
		block.number,
		block.prevHash,
		// Sha of uncles
		"",
		block.coinbase,
		// root state
		"",
		string(Sha256Bin([]byte(Encode(encTx)))),
		block.difficulty,
		uint64(block.time),
		block.nonce,
		// extra stuff?
	}

	uncles := []interface{}{}

	return Encode([]interface{}{header, encTx, uncles})
}

func (block *Block) UnmarshalRlp(data []byte) {
	t, _ := Decode(data, 0)
	if slice, ok := t.([]interface{}); ok {
		if header, ok := slice[0].([]interface{}); ok {
			if number, ok := header[0].(uint8); ok {
				block.number = uint32(number)
			}

			if prevHash, ok := header[1].([]byte); ok {
				block.prevHash = string(prevHash)
			}

			if coinbase, ok := header[3].([]byte); ok {
				block.coinbase = string(coinbase)
			}

			if difficulty, ok := header[6].(uint8); ok {
				block.difficulty = uint32(difficulty)
			}
			if difficulty, ok := header[6].(uint64); ok {
				block.difficulty = uint32(difficulty)
			}

			if time, ok := header[7].(uint8); ok {
				block.time = int64(time)
			}
			if time, ok := header[7].(uint64); ok {
				block.time = int64(time)
			}

			if nonce, ok := header[8].(uint8); ok {
				block.nonce = uint32(nonce)
			}

			if extra, ok := header[9].([]byte); ok {
				block.extra = string(extra)
			}
		}

		if txSlice, ok := slice[1].([]interface{}); ok {
			block.transactions = make([]*Transaction, len(txSlice))
			for i, tx := range txSlice {
				if t, ok := tx.([]byte); ok {
					tx := &Transaction{}
					tx.UnmarshalRlp(t)
					block.transactions[i] = tx
				}
			}
		}
	}
}
