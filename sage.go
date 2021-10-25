package main

import "flag"

const Debug = false

var StartDBQueryInterface bool

func Init() {
	flag.BoolVar(&StartDBQueryInterface, "db", false, "start db query interface")
	flag.Parse()
}

// func RegisterInterrupts(s *Server) {

// }

func main() {
	InitFees()
	Init()

	if StartDBQueryInterface {
		dbInterface := NewDBInterface()
		dbInterface.Start()
	} else {
		Testing()
	}
}
