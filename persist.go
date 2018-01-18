package Juicy

import (
	"fmt"

	rbt "github.com/salamer/RbTree"
)

func (db *DB) Persist(filename string) {
	node, err := db.Tree.GetRoot()
	if err != nil {
		fmt.Println(err)
	} else {
		db.serialize(filename, node)
	}
}

func (db *DB) serialize(filename string, node *rbt.Node) {
	fmt.Println(node)
	if node.Left != nil {
		db.serialize(filename, node.Left)
	}
	if node.Right != nil {
		db.serialize(filename, node.Right)
	}
}
