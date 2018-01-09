package Juicy

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type DB struct {
	tree *rbt.Tree
	name string
}

type Node struct {
	next  *node
	key   string
	value interface{}
}

func NewDB(name string) *DB {
	return &DB{
		tree: rbt.NewWithIntComparator(),
		name: name,
	}
}

func NewNode(key string, value interface{}) *Node {
	return &Node{
		key:   string,
		value: value,
	}
}

func (db *DB) GetValue(key string) (interface{}, error) {
	node, r = db.tree.Get(Hash(key))
	if !r {
		return nil, KeyError
	} else {
		for node != nil {
			if node.key == key {
				return node.value, nil
			}
			node = node.next
		}
	}
}

func (db *DB) SetValue(key string, value interface{}) error {
	k := Hash(key)
	if node, r := db.GetValue(k); r {
		haveKey := false
		for node != nil {
			// for the same key
			if node.key == key {
				node.value = value
				haveKey = true
			}
			node = node.next
		}
		node.next = NewNode(k, value)
	} else {
		node = NewNode(k, value)
		db.tree.Put(k, node)
	}
}
