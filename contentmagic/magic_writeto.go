package okatype_contentmagic

import (
	"fmt"
	"io"
)

// WriteTo makes okatype_contentmagic.Type fit the io.WriterTo interface.
func (receiver Type) WriteTo(w io.Writer) (int64, error) {
	const doesNotMatter = int64(-1)

	if nil == w {
		return doesNotMatter, errNilWriter
	}

	// We do this (i.e., copy the values over to a local buffer) to prevent the possibility of `w.Write()` writing to the `magic` global.
	//
	// We are OK with using the constant "8" here, because it is a compile time error if the length of `buffer` and the `magic` global don't match.
	var buffer [8]byte = magic

	p := buffer[:]
	n, err := w.Write(p)
	if nil != err {
		return doesNotMatter, err
	}

	if expected, actual := len(magic), n; expected != actual {
		var n64 int64 = int64(n)
		if expected, actual := n, int(n64); expected != actual { // Future proof, check for overflow.
			n64 = doesNotMatter
		}

		return n64, io.ErrShortWrite
	}

	var n64 int64 = int64(n)
	if expected, actual := n, int(n64); expected != actual { // Future proof, check for overflow.
		return doesNotMatter, fmt.Errorf("int64 overflow: cannot convert (%T) %#v into and int64.", expected, expected)
	}

	return n64, nil
}
