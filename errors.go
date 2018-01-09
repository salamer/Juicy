package Juicy

import "errors"

var (
	KeyError   = errors.New("key not in database")
	ValueError = errors.New("Value Error")
)
