package main

import "math/big"

// BigPow returns big.Int
func BigPow(a, b int) *big.Int {
	c := new(big.Int)
	c.Exp(
		big.NewInt(int64(a)),
		big.NewInt(int64(b)),
		big.NewInt(0),
	)
	return c
}

// Big returns big.Int
func Big(num string) *big.Int {
	n := new(big.Int)
	n.SetString(num, 0)

	return n
}
