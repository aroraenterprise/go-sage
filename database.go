package main

import (
	"fmt"
	"os/user"
	"path"

	"github.com/syndtr/goleveldb/leveldb"
)

type LDBDatabase struct {
	db   *leveldb.DB
	trie *Trie
}

func NewDatabase() (*LDBDatabase, error) {
	usr, _ := user.Current()
	dbPath := path.Join(usr.HomeDir, ".sage", "database")

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, err
	}
	database := &LDBDatabase{
		db: db,
	}

	err = database.Bootstrap()
	return database, err
}

func (db *LDBDatabase) Bootstrap() error {
	db.trie = NewTrie(db)
	return nil
}

func (db *LDBDatabase) Put(key []byte, value []byte) {
	err := db.db.Put(key, value, nil)
	if err != nil {
		fmt.Println("Error put", err)
	}
}

func (db *LDBDatabase) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (db *LDBDatabase) Close() {
	db.db.Close()
}
