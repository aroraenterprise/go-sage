package main

import (
	"math"
)

/*
 * This is the special genesis block.
 */

var GenesisHeader = []interface{}{
	// Block number
	uint32(0),
	// Previous hash (none)
	"",
	// Sha of uncles
	string(Sha256Bin(Encode([]interface{}{}))),
	// Coinbase
	"",
	// Root state
	"",
	// Sha of transactions
	string(Sha256Bin(Encode([]interface{}{}))),
	// Difficulty
	uint32(math.Pow(2, 36)),
	// Time
	uint64(1),
	// Nonce
	uint32(0),
	// Extra
	"",
}

var Genesis = []interface{}{GenesisHeader, []interface{}{}, []interface{}{}}

var GenesisBlock = NewBlock(Encode(Genesis))
