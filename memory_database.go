package main

type MemDatabase struct {
	db   map[string][]byte
	trie *Trie
}

func NewMemDatabase() (*MemDatabase, error) {
	db := &MemDatabase{
		db: make(map[string][]byte),
	}
	db.trie = NewTrie(db, "")
	return db, nil
}

func (db *MemDatabase) Put(key []byte, value []byte) {
	db.db[string(key)] = value
}

func (db *MemDatabase) Get(key []byte) ([]byte, error) {
	return db.db[string(key)], nil
}
