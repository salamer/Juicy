package main

import (
	"fmt"

	client "github.com/salamer/Juicy/client"
)

func main() {
	c := client.NewJuicyClient("localhost", 8080)
	c.Set("helloq", "world")
	c.Set("hahah", "wwww")
	c.Set("aaaa", "zzzz")
	fmt.Println(c.Get("aaaa"))
	fmt.Println(c.Get("helloq"))
	fmt.Println(c.Get("hahah"))
	fmt.Println(c.Delete("hahah"))
	fmt.Println(c.Get("hahah"))
	c.Persist("aaa.txt")
	fmt.Println(c.Empty())
	c.Clear()
	fmt.Println(c.Empty())
}
