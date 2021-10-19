package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

// Op codes
const (
	oSTOP      int = 0x00
	oADD       int = 0x10
	oSUB       int = 0x11
	oMUL       int = 0x12
	oDIV       int = 0x13
	oSDIV      int = 0x14
	oMOD       int = 0x15
	oSMOD      int = 0x16
	oEXP       int = 0x17
	oNEG       int = 0x18
	oLT        int = 0x20
	oLE        int = 0x21
	oGT        int = 0x22
	oGE        int = 0x23
	oEQ        int = 0x24
	oNOT       int = 0x25
	oSHA256    int = 0x30
	oRIPEMD160 int = 0x31
	oECMUL     int = 0x32
	oECADD     int = 0x33
	oSIGN      int = 0x34
	oRECOVER   int = 0x35
	oCOPY      int = 0x40
	oST        int = 0x41
	oLD        int = 0x42
	oSET       int = 0x43
	oJMP       int = 0x50
	oJMPI      int = 0x51
	oIND       int = 0x52
	oEXTRO     int = 0x60
	oBALANCE   int = 0x61
	oMKTX      int = 0x70
	oDATA      int = 0x80
	oDATAN     int = 0x81
	oMYADDRESS int = 0x90
	oSUICIDE   int = 0xff
)

type OpType int

const (
	tNorm = iota
	tData
	tExtro
	tCrypto
)

type TxCallback func(opType OpType) bool

type Vm struct {
	stack  map[string]string
	iptr   int
	memory map[string]map[string]string
}

func NewVm() *Vm {
	fmt.Println("init sage vm")
	stackSize := uint(256)
	fmt.Println("stack size =", stackSize)

	return &Vm{make(map[string]string), 0, make(map[string]map[string]string)}
}

func (vm *Vm) RunTransaction(tx *Transaction, cb TxCallback) {
	fmt.Printf("# processing Tx (%v)\n", tx.addr)
	fmt.Printf("fee = %f\n", float32(tx.fee)/1e8)
	fmt.Printf("ops = %d\n", len(tx.data))
	fmt.Printf("sender = %s\n", tx.sender)
	fmt.Printf("value = %d\n", tx.value)

	vm.stack = make(map[string]string)
	vm.stack["0"] = tx.sender
	vm.stack["1"] = "100"
	vm.memory[tx.addr] = make(map[string]string)

	x := 0
	y := 1
	z := 2 //a := 3; b := 4; c := 5
out:
	for vm.iptr < len(tx.data) {

		base := new(big.Int)
		op, args, _ := Instr(tx.data[vm.iptr])
		fmt.Printf("%-3d %d %v\n", vm.iptr, op, args)

		opType := OpType(tNorm)
		switch op {
		case oEXTRO, oBALANCE:
			opType = tExtro

		case oSHA256, oRIPEMD160, oECMUL, oECADD:
			opType = tCrypto
		}

		if !cb(opType) {
			break out
		}

		nptr := vm.iptr
		switch op {
		case oSTOP:
			fmt.Println("exiting (oSTOP), idx =", nptr)

			break out
		case oADD:
			// (Rx + Ry) % 2 ** 256
			base.Add(Big(vm.stack[args[x]]), Big(vm.stack[args[y]]))
			base.Mod(base, big.NewInt(int64(math.Pow(2, 256))))
			// Set the result to Rz
			vm.stack[args[z]] = base.String()
		case oSUB:
			// (Rx - Ry) % 2 ** 256
			base.Sub(Big(vm.stack[args[x]]), Big(vm.stack[args[y]]))
			base.Mod(base, big.NewInt(int64(math.Pow(2, 256))))
			// Set the result to Rz
			vm.stack[args[z]] = base.String()
		case oMUL:
			// (Rx * Ry) % 2 ** 256
			base.Mul(Big(vm.stack[args[x]]), Big(vm.stack[args[y]]))
			base.Mod(base, big.NewInt(int64(math.Pow(2, 256))))
			// Set the result to Rz
			vm.stack[args[z]] = base.String()
		case oDIV:
			// floor(Rx / Ry)
			base.Div(Big(vm.stack[args[x]]), Big(vm.stack[args[y]]))
			// Set the result to Rz
			vm.stack[args[z]] = base.String()
		case oSET:
			// Set the (numeric) value at Iy to Rx
			vm.stack[args[x]] = args[y]
		case oLD:
			// Load the value at Mx to Ry
			vm.stack[args[y]] = vm.memory[tx.addr][vm.stack[args[x]]]
		case oLT:
			cmp := Big(vm.stack[args[x]]).Cmp(Big(vm.stack[args[y]]))
			// Set the result as "boolean" value to Rz
			if cmp < 0 { // a < b
				vm.stack[args[z]] = "1"
			} else {
				vm.stack[args[z]] = "0"
			}
		case oJMP:
			// Set the instruction pointer to the value at Rx
			ptr, _ := strconv.Atoi(vm.stack[args[x]])
			nptr = ptr
		case oJMPI:
			// Set the instruction pointer to the value at Ry if Rx yields true
			if vm.stack[args[x]] != "0" {
				ptr, _ := strconv.Atoi(vm.stack[args[y]])
				nptr = ptr
			}
		default:
			fmt.Println("Error op", op)
			break
		}

		if vm.iptr == nptr {
			vm.iptr++
		} else {
			vm.iptr = nptr
			fmt.Println("... JMP", nptr, "...")
		}
	}

	fmt.Println("# finished processsing Tx")
}
