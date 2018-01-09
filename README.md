
#  Juicy :cocktail:

> "It was all a dream, I used to read Word Up! magazine."
>
> -- <Juicy> The Notorious B.I.G.

# INSTALL

    go get -u github.com/salamer/Juicy

# QUICK START

### Single Node

```GO

package main

import (
	"fmt"

	Juicy "github.com/salamer/Juicy"
)

func main() {
	db := Juicy.NewDB("hello", Juicy.SINGLE, Juicy.RaftConf{})
	db.Start()
	db.SetValue("hello", "world")
	db.SetValue("lalala", "zzzz")
	db.SetValue("oh", "haha")
	fmt.Println(db.GetValue("oh"))
}


```

### Multi Node

```GO

package main

import (
	"fmt"

	Juicy "github.com/salamer/Juicy"
)

const ConfPath = {confpath}      //your node conf json file
const ID = {id}                  // your node id
const Name = {name}              //your node name
const Host = {host}              //your raft node host
const Port = {port}              //your raft node port

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



```

## LICENSE
Copyright © 2018 by Aljun

Under Apache license : [http://www.apache.org/licenses/](http://www.apache.org/licenses/)

