package main

import (
	"fmt"

	Juicy "github.com/salamer/Juicy"
)

const ConfPath = ""      //your node conf json file
const ID = 1             // your node id
const Name = "aljun"     //your node name
const Host = "localhost" //your raft node host
const Port = 8000        //your raft node port

func main() {
	db := Juicy.NewDB("hello", Juicy.DISTRIBUTED, Juicy.RaftConf{
		ID:       ID,
		Name:     Name,
		ConfPath: ConfPath,
		Port:     Port,
		Host:     Host,
	})
	db.Start()
	db.SetValue("hello", "world")
	db.SetValue("lalala", "zzzz")
	db.SetValue("oh", "haha")
	fmt.Println(db.GetValue("oh"))
}
