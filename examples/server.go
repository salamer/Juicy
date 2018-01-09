package main

import (
	"fmt"

	Juicy "github.com/salamer/Juicy"
)

func main() {
	db := Juicy.NewDB("hello", Juicy.SINGLE)
	db.SetValue("hello", "world")
	db.SetValue("lalala", "zzzz")
	db.SetValue("oh", "haha")
	fmt.Println(db.GetValue("oh"))
}
