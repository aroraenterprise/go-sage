package main

import "fmt"

type Database interface {
	Put(key []byte, value []byte)
	Get(key []byte) ([]byte, error)
}

type Trie struct {
	root string
	db   Database
}

func NewTrie(db Database) *Trie {
	return &Trie{db: db, root: ""}
}

func (t *Trie) Update(key, value string) {
	k := CompactHexDecode(key)
	t.root = t.UpdateState(t.root, k, value)
}

func (t *Trie) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (t *Trie) Put(node []byte) []byte {
	enc := Encode(node)
	fmt.Println("data", enc)
	sha := Sha256Bin(enc)

	t.db.Put(sha, enc)
	return sha
}

func (t *Trie) UpdateState(node string, key []int, value string) string {
	if value == "" {
		return t.InsertState(node, key, value)
	}
	return ""
}

func (t *Trie) InsertState(node string, key []int, value string) string {
	return ""
}
