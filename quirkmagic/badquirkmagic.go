package okatype_quirkmagic

import (
	"bytes"
	"fmt"
)

// BadQuirkMagic represents a "bad quirk magic" error.
//
// You can detect if an error is a "bad quirk magic" error, with code like the following:
//
//	n64, err := quirkmagic.ReadFrom(r)
//	
//	switch err.(type) {
//	case okatype_quirkmagic.BadQuirkMagic:
//		//@TODO
//	default:
//		//@TODO
//	}
type BadQuirkMagic interface {
	error
	BadQuirkMagic()

	// Reader returns a Reader that can be used to get the "bad quirk magic" that what was read in.
	//
	// So, for example, what would have been expected to have been read in would have been:
	//
	//	[8]byte{7,'o','k','a','n','e','r','o'}
	//
	// But what was actually instead read in was:
	//
	//	[8]byte{'<','?','x','m','l',' ','v','e'}
	//
	// This might be useful if you want to create an informative error message.
	Reader() Reader
}

type internalBadQuirkMagic struct {
	msg string
	expected []byte
	actual []byte
}

func newBadQuirkMagic(expected []byte, actual []byte) BadQuirkMagic {
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

	msg := fmt.Sprintf("bad quirk magic: expected (%d, %q) [len=%d], but actually got (%d, %q) [len=%d].", e0, e, len(copyExpected), a0, a, len(copyActual))

	complainer := internalBadQuirkMagic{
		msg:msg,
		expected: copyExpected,
		actual:   copyActual,
	}

	return &complainer
}

func (receiver internalBadQuirkMagic) Error() string {
	return receiver.msg
}

func (internalBadQuirkMagic) BadQuirkMagic() {
	// Nothing here.
}

func (receiver internalBadQuirkMagic) Reader() Reader {
	return bytes.NewReader(receiver.actual)
}
