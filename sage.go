package main

func main() {
	InitFees()
	bm := NewBlockManager()
	tx := NewTransaction(0x0, 20, []string{
		"SET 10 6",
		"LD 10 10",
		"LT 10 1 20",
		"SET 255 7",
		"STOP",
		"SET 30 200",
		"LD 30 31",
		"SET 255 22",
		"SET 255 15",
	})

	tx2 := NewTransaction(0x0, 20, []string{"SET 10 6", "LD 10 10"})
	blck := NewBlock([]*Transaction{tx2, tx})
	bm.ProcessBlock(blck)
}
