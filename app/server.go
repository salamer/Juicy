package main

import (
	Juicy "github.com/salamer/Juicy"
)

func main() {
	db := Juicy.NewDB("hello", Juicy.SINGLE, Juicy.RaftConf{})
	db.Start()
}
