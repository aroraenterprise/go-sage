package main

import "time"

type Block struct {
	RlpSerializer
	number       uint32
	prevHash     string
	uncles       []*Block
	coinbase     string
	difficulty   int
	time         time.Time
	nonce        int
	transactions []*Transaction
}

func NewBlock(transactions []*Transaction) *Block {
	block := &Block{
		transactions: transactions,
		time:         time.Now(),
	}

	return block
}

func (block *Block) Hash() string {
	return Sha256Hex(block.MarshalRlp())
}

func (block *Block) MarshalRlp() []byte {
	encTx := make([]string, len(block.transactions))
	for i, tx := range block.transactions {
		encTx[i] = string(tx.MarshalRlp())
	}

	enc := RlpEncode([]interface{}{
		block.number,
		block.prevHash,
		// Sha of uncles
		block.coinbase,
		// root state
		Sha256Bin([]byte(RlpEncode(encTx))),
		block.difficulty,
		block.time,
		block.nonce,
		// extra stuff?
	})

	return []byte(enc)
}

func (block *Block) UnmarshalRlp(data []byte) {

}
