package Juicy

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

const (
	SINGLE = iota
	DISTRIBUTED
)

type DB struct {
	Tree *rbt.Tree
	name string
}

type Node struct {
	next  *Node
	key   string
	value interface{}
}

func NewDB(name string, mode int) *DB {
	return &DB{
		Tree: rbt.NewWithIntComparator(),
		name: name,
	}
}

func NewNode(key string, value interface{}) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func (db *DB) GetValue(key string) (interface{}, error) {
	node, r := SafeString(db.Tree.Get(Hash(key)))
	if r != nil {
		return nil, KeyError
	} else {
		for node != nil {
			if node.key == key {
				return node.value, nil
			}
			node = node.next
		}
		return node.value, nil
	}
}

func (db *DB) GetNode(key string) (*Node, error) {
	node, r := SafeString(db.Tree.Get(Hash(key)))
	if r != nil {
		return nil, KeyError
	} else {
		return node, nil
	}
}

func (db *DB) SetValue(key string, value interface{}) error {
	node, r := db.GetNode(key)
	if r != nil {
		node = NewNode(key, value)
		db.Tree.Put(Hash(key), node)
		return nil
	} else {
		haveKey := false
		for node.next != nil {
			// for the same key
			if node.key == key {
				node.value = value
				haveKey = true
			}
			node = node.next
		}
		if !haveKey {
			node.next = NewNode(key, value)
		}
		return nil
	}

}
