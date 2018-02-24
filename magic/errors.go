package okatype_magic

import (
	"errors"
)

var (
	errNilReader   = errors.New("Nil Reader")
	errNilReceiver = errors.New("Nil Receiver")
	errNilWriter   = errors.New("Nil Writer")
)
