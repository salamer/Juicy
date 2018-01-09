
#  Juicy :cocktail:

> "It was all a dream, I used to read Word Up! magazine."
>
> -- <Juicy> The Notorious B.I.G.

# INSTALL

    go get -u github.com/salamer/Juicy

# QUICK START

### Single Node

```

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

```

## LICENSE
Copyright © 2018 by Aljun

Under Apache license : [http://www.apache.org/licenses/](http://www.apache.org/licenses/)

