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

func NewTrie(db Database, root string) *Trie {
	return &Trie{db: db, root: ""}
}

func (t *Trie) Update(key, value string) {
	k := CompactHexDecode(key)
	t.root = t.UpdateState(t.root, k, value)
}

func (t *Trie) Get(key string) ([]byte, error) {
	return nil, nil
}

func (t *Trie) Put(node interface{}) []byte {
	enc := Encode(node)
	sha := Sha256Bin(enc)

	t.db.Put([]byte(sha), enc)
	return sha
}

func (t *Trie) UpdateState(node string, key []int, value string) string {
	if value != "" {
		return t.InsertState(node, key, value)
	}
	return ""
}

func DecodeNode(data []byte) []string {
	dec, _ := Decode(data, 0)
	if slice, ok := dec.([]interface{}); ok {
		strSlice := make([]string, len(slice))

		for i, s := range slice {
			if str, ok := s.([]byte); ok {
				strSlice[i] = string(str)
			}
		}

		return strSlice
	}

	return nil
}

func (t *Trie) InsertState(node string, key []int, value string) string {
	if len(key) == 0 {
		return value
	}
	if node == "" {
		newNode := []string{CompactEncode(key), value}
		return string(t.Put(newNode))
	}
	n, err := t.db.Get([]byte(node))

	if err != nil {
		fmt.Println("Error InsertState", err)
		return ""
	}

	currentNode := DecodeNode(n)
	if len(currentNode) == 2 {
		k := CompactDecode(currentNode[0])
		v := currentNode[1]

		if CompareIntSlice(k, key) {
			return string(t.Put([]string{CompactEncode(key), value}))
		}
		var newHash string
		matchingLength := MatchingNibbleLength(key, k)
		if matchingLength == len(k) {
			// insert the hash, create new node
			newHash = t.InsertState(v, key[matchingLength:], value)
		} else {
			oldNode := t.InsertState("", k[matchingLength+1:], v)
			newNode := t.InsertState("", key[matchingLength+1:], value)

			scaledSlice := make([]string, 17)
			scaledSlice[k[matchingLength]] = oldNode
			scaledSlice[key[matchingLength]] = newNode
			newHash = string(t.Put(scaledSlice))
		}

		if matchingLength == 0 {
			return newHash
		} else {
			newNode := []string{CompactEncode(key[:matchingLength]), newHash}
			return string(t.Put(newNode))
		}
	} else {
		newNode := make([]string, 17)
		copy(newNode, currentNode)
		newNode[key[0]] = t.InsertState(currentNode[key[0]], key[1:], value)
		return string(t.Put(newNode))
	}
}
