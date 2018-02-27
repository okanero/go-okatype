package okatype_pascalstring

import (
	"errors"
)

var (
	errNilWriter = errors.New("Nil Writer")
	errTooLong   = errors.New("Too Long")
)
