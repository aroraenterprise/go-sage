package main

import (
	"encoding/hex"
	"testing"
)

var testsource = `{"Inputs": {
		"doe": "reindeer",
		"dog": "puppy",
		"dogglesworth": "cat"
	},
	"Expectation": "e378927bfc1bd4f01a2e8d9f59bd18db8a208bb493ac0b00f93ce51d4d2af76c"
}`

func TestTestRunner(t *testing.T) {
	db, _ := NewMemDatabase()
	trie := NewTrie(db, "")

	runner := NewTestRunner(t)
	runner.RunFromString(testsource, func(source *TestSource) {
		for key, value := range source.Inputs {
			trie.Update(key, value)
		}
		if hex.EncodeToString([]byte(trie.root)) != source.Expectation {
			t.Error("trie root did not match")
		}
	})
}
