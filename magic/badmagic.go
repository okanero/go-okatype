package okatype_magic

import (
	"bytes"
	"fmt"
)

// BadMagic represents a "bad magic" error.
//
// You can detect if an error is a "bad magic" error, with code like the following:
//
//	n64, err := magic.ReadFrom(r)
//	
//	switch err.(type) {
//	case okatype_magic.BadMagic:
//		//@TODO
//	default:
//		//@TODO
//	}
type BadMagic interface {
	error
	BadMagic()
	Reader() Reader
}

type internalBadMagic struct {
	msg string
	expected []byte
	actual []byte
}

func newBadMagic(expected []byte, actual []byte) BadMagic {
	copyExpected := append([]byte(nil), expected...)
	copyActual   := append([]byte(nil), actual...)

	var e0 byte
	var a0 byte

	var e []byte
	var a []byte

	if 1 <= len(copyExpected) {
		e0 = copyExpected[0]
		e  = copyExpected[1:]
	}
	if 1 <= len(copyActual) {
		a0 = copyActual[0]
		a  = copyActual[1:]
	}

	msg := fmt.Sprintf("bad magic: expected (%d, %q) [len=%d], but actually got (%d, %q) [len=%d].", e0, e, len(copyExpected), a0, a, len(copyActual))

	complainer := internalBadMagic{
		msg:msg,
		expected: copyExpected,
		actual:   copyActual,
	}

	return &complainer
}

func (receiver internalBadMagic) Error() string {
	return receiver.msg
}

func (internalBadMagic) BadMagic() {
	// Nothing here.
}

func (receiver internalBadMagic) Reader() Reader {
	return bytes.NewReader(receiver.actual)
}
