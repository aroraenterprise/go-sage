package main

import (
	"encoding/hex"
	"testing"
)

func TestTriePut(t *testing.T) {
	db, err := NewMemDatabase()

	if err != nil {
		t.Error("Failed to create mem database")
	}

	key := db.trie.Put([]byte("testing node"))

	data, err := db.Get(key)
	if err != nil {
		t.Error("Nothing at node")
	}

	s, _ := Decode(data, 0)
	if str, ok := s.([]byte); ok {
		if string(str) != "testing node" {
			t.Error("Wrong value at node", str)
		}
	} else {
		t.Error("Wrong return type")
	}

}

func TestTrieUpdate(t *testing.T) {
	db, err := NewMemDatabase()
	trie := NewTrie(db, "")

	if err != nil {
		t.Error("Error starting db")
	}

	trie.Update("doe", "reindeer")
	trie.Update("dog", "puppy")
	trie.Update("dogglesworth", "cat")

	root := hex.EncodeToString([]byte(trie.root))
	req := "e378927bfc1bd4f01a2e8d9f59bd18db8a208bb493ac0b00f93ce51d4d2af76c"

	if root != req {
		t.Error("trie.root do not match, expected", req, "got", root)
	}

	db.trie.Update("test", "test")
}
