package Juicy

import "hash/fnv"

func Hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

//get the key string

func SafeString(a interface{}, b bool) (*Node, error) {
	if !b {
		return nil, KeyError
	} else {
		if val, ok := a.(*Node); ok {
			return val, nil

		}
		return nil, ValueError
	}
}
