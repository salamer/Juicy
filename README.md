
#  Juicy :cocktail:

> "It was all a dream, I used to read Word Up! magazine."
>
> -- <Juicy> The Notorious B.I.G.

# INSTALL

    go get -u github.com/salamer/Juicy

# QUICK START

### USEING PYTHON CLIENT

> pip install Juicy

```Python

>>> from Juicy import Client
>>> j = Client("localhost",8000)
>>> j.set("hello","world")
success: true

>>> j.get("hello")
success: true
res2: "world"

>>> j.set("aljun","coder")
success: true

>>> "aljun" in j
True
>>> j.delete("aljun")
success: true

>>> j.get("aljun")
error: "key not in database"

```

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

