package main

type Block struct {
	transactions []*Transaction
}

func NewBlock(transactions []*Transaction) *Block {
	block := &Block{
		transactions: transactions,
	}

	return block
}
