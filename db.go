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
	size int
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

func GetDBFromFile(filename string) *DB {
	//TODO: read from persist file
	return &DB{}
}

func NewDB(name string, mode int, conf RaftConf) *DB {
	if mode == SINGLE {
		return &DB{
			Tree: rbt.NewWithIntComparator(),
			name: name,
			raft: nil,
			Mode: SINGLE,
			size: 0,
		}
	} else {
		return &DB{
			Tree: rbt.NewWithIntComparator(),
			name: name,
			raft: raft.NewNode(conf.Name, conf.ID, conf.Host, conf.Port, conf.ConfPath),
			Mode: DISTRIBUTED,
			size: 0,
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
		db.size += 1
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
		db.size += 1
		return nil
	}

}

func (db *DB) Delete(key string) error {
	node, r := db.GetNode(key)
	if r != nil {
		return r
	} else {
		if node.key == key {
			db.Tree.Remove(Hash(node.key))
		} else {
			_node := node.next
			for _node != nil {
				if _node.key == key {
					node.next = _node.next
					return nil
				}
				_node = _node.next
				node = node.next
			}
		}
		db.size -= 1
		return nil
	}
}

func (db *DB) Size() int {
	return db.size
}

func (db *DB) HaveKey(key string) (bool, error) {
	node, r := db.GetNode(key)
	if r != nil && node != nil {
		return false, KeyError
	} else {
		return true, nil
	}
}

func (db *DB) Clear() {
	db.Tree.Clear()
}

func (db *DB) Empty() bool {
	if db.Tree.Empty() {
		return true
	} else {
		return false
	}
}
