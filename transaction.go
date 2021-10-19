package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

var StepFee *big.Int = new(big.Int)
var TxFee *big.Int = new(big.Int)
var MemFee *big.Int = new(big.Int)
var DataFee *big.Int = new(big.Int)
var CryptoFee *big.Int = new(big.Int)
var ExtroFee *big.Int = new(big.Int)

var Period1Reward *big.Int = new(big.Int)
var Period2Reward *big.Int = new(big.Int)
var Period3Reward *big.Int = new(big.Int)
var Period4Reward *big.Int = new(big.Int)

type Transaction struct {
	sender    string
	recipient uint32
	value     uint32
	fee       uint32
	data      []string
	memory    []string
	signature string
	addr      string
}

// NewTransaction returns Transaction
func NewTransaction(to uint32, value uint32, data []string) *Transaction {
	tx := Transaction{sender: "12", recipient: to, value: value}
	tx.fee = 0

	tx.data = make([]string, len(data))
	for i, val := range data {
		instr, err := CompileInstr(val)
		if err != nil {
			fmt.Printf("compile error:%d %v", i+1, err)
		}
		tx.data[i] = instr
	}

	b := []byte(tx.Serialize())
	hash := sha256.Sum256(b)
	tx.addr = hex.EncodeToString(hash[:])
	return &tx
}

func Uitoa(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func (tx *Transaction) Serialize() string {
	preEnc := []interface{}{
		"0", // TODO last tx
		tx.sender,
		Uitoa(tx.recipient),
		Uitoa(tx.value),
		Uitoa(tx.fee),
		tx.data,
	}
	return RlpEncode(preEnc)
}

func InitFees() {
	b60 := new(big.Int)
	b60.Exp(big.NewInt(2), big.NewInt(60), big.NewInt(0))

	b80 := new(big.Int)
	b80.Exp(big.NewInt(2), big.NewInt(80), big.NewInt(0))

	StepFee.Mul(b60, big.NewInt(4096))
	TxFee.Mul(b60, big.NewInt(524288))
	MemFee.Mul(b60, big.NewInt(262144))
	DataFee.Mul(b60, big.NewInt(16384))
	CryptoFee.Mul(b60, big.NewInt(65536))
	ExtroFee.Mul(b60, big.NewInt(65536))

	Period1Reward.Mul(b80, big.NewInt(1024))
	//fmt.Println("Period1Reward:", Period1Reward)

	Period2Reward.Mul(b80, big.NewInt(512))
	//fmt.Println("Period2Reward:", Period2Reward)

	Period3Reward.Mul(b80, big.NewInt(256))
	//fmt.Println("Period3Reward:", Period3Reward)

	Period4Reward.Mul(b80, big.NewInt(128))
	//fmt.Println("Period4Reward:", Period4Reward)
}
