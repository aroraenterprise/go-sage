package main

type BlockManager struct {
	vm *Vm
}

func NewBlockManager() *BlockManager {
	bm := &BlockManager{vm: NewVm()}

	return bm
}

func (bm *BlockManager) ProcessBlock(block *Block) error {
	txCount := len(block.transactions)
	lockChan := make(chan bool, txCount)

	for _, tx := range block.transactions {
		go bm.ProcessTransaction(tx, lockChan)
	}
	for i := 0; i < txCount; i++ {
		<-lockChan
	}
	return nil
}

func (bm *BlockManager) ProcessTransaction(tx *Transaction, lockChan chan bool) {
	if tx.recipient == "\x00" {
		bm.vm.RunTransaction(tx, func(opType OpType) bool {
			// TODO calculate fees
			return true
		})
	}

	lockChan <- true
}
