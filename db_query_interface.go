package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

type DbInterface struct {
	db   *MemDatabase
	trie *Trie
}

func NewDBInterface() *DbInterface {
	db, _ := NewMemDatabase()
	trie := NewTrie(db, "")

	return &DbInterface{db: db, trie: trie}
}

func (i *DbInterface) ValidateInput(action string, argumentLength int) error {
	err := false
	var expArgCount int

	switch {
	case action == "update" && argumentLength != 2:
		err = true
		expArgCount = 2
	case action == "get" && argumentLength != 1:
		err = true
		expArgCount = 1
	case (action == "quit" || action == "exit") && argumentLength != 0:
		err = true
		expArgCount = 0
	}

	if err {
		return errors.New(fmt.Sprintf("'%s' requires %d args, got %d", action, expArgCount, argumentLength))
	} else {
		return nil
	}
}

func (i *DbInterface) ParseInput(input string) bool {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	count := 0
	var tokens []string
	for scanner.Scan() {
		count++
		tokens = append(tokens, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	if len(tokens) == 0 {
		return true
	}

	err := i.ValidateInput(tokens[0], count-1)
	if err != nil {
		fmt.Println(err)
	} else {
		switch tokens[0] {
		case "update":
			i.trie.Update(tokens[1], tokens[2])
			fmt.Println(hex.EncodeToString([]byte(i.trie.root)))

		case "get":
			fmt.Println(i.trie.Get(tokens[1]))

		case "exit", "quit", "q":
			return false

		default:
			fmt.Println("Unknown command:", tokens[0])
		}
	}
	return true
}

func (i *DbInterface) Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("db >>>")
		str, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Error reading input", err)
		} else {
			if !i.ParseInput(string(str)) {
				return
			}
		}
	}
}
