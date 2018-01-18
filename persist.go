package Juicy

import (
	"log"
	"os"
	"strings"

	rbt "github.com/salamer/RbTree"
)

const SEPARATOR = " "
const SEPARATOR_PLACEHOLER = "\t"
const Newline = "\n"

func (db *DB) Persist(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Println("create file err: ", err)
		}
		defer file.Close()
	}
	node, err := db.Tree.GetRoot()
	if err != nil {
		log.Println(err)
	} else {
		db.serialize(filename, node)
	}
}

func (db *DB) serialize(filename string, node *rbt.Node) {
	if _node, ok := node.Value.(*Node); ok {
		data := []byte(SEPARATOR)
		for _node != nil {
			// replace space in key
			value := strings.Replace(_node.value, SEPARATOR_PLACEHOLER, SEPARATOR_PLACEHOLER+SEPARATOR_PLACEHOLER, -1)
			value = strings.Replace(value, SEPARATOR, SEPARATOR_PLACEHOLER, -1)

			_value := []byte(value)
			_key := []byte(_node.key)

			// the format is like "{key} {value}"
			data = append(data[:], _key[:]...)
			data = append(data[:], []byte(SEPARATOR)...)
			data = append(data[:], _value[:]...)

			_node = _node.next
		}
		data = append(data[:], []byte(Newline)...)
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println("open file err:", err)
		}
		_, err = f.Write(data)
		if err != nil {
			log.Println("write file err:", err)
		}
		f.Close()

	}
	if node.Left != nil {
		db.serialize(filename, node.Left)
	}
	if node.Right != nil {
		db.serialize(filename, node.Right)
	}
}
