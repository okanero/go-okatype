package okatype_quirkmagic

import (
	"io"
)

// Reader is an interface that wraps a full suite of reading functionality.
//
// The methods in this interface are meant to match the methods in the bytes.Reader struct.
type Reader interface {
	Len() int
	Read(b []byte) (n int, err error)
	ReadAt(b []byte, off int64) (n int, err error)
	ReadByte() (byte, error)
	ReadRune() (ch rune, size int, err error)
	Reset(b []byte)
	Seek(offset int64, whence int) (int64, error)
	Size() int64
	UnreadByte() error
	UnreadRune() error
	WriteTo(w io.Writer) (n int64, err error)
}
