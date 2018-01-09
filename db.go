package Juicy

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	raft "github.com/salamer/naive_raft"
)

const (
	SINGLE = iota
	DISTRIBUTED
)

type DB struct {
	Tree *rbt.Tree
	name string
	raft *raft.Node
	Mode int
}

type Node struct {
	next  *Node
	key   string
	value interface{}
}

type RaftConf struct {
	ID       int
	Name     string
	ConfPath string
	Port     int
	Host     string
}

func NewDB(name string, mode int, conf RaftConf) *DB {
	if mode == SINGLE {
		return &DB{
			Tree: rbt.NewWithIntComparator(),
			name: name,
			raft: nil,
			Mode: SINGLE,
		}
	} else {
		return &DB{
			Tree: rbt.NewWithIntComparator(),
			name: name,
			raft: raft.NewNode(conf.Name, conf.ID, conf.Host, conf.Port, conf.ConfPath),
			Mode: DISTRIBUTED,
		}
	}
}

func (db *DB) Start() error {
	if db.Mode == SINGLE {
		return nil
	} else {
		db.raft.Run()
		return nil
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
