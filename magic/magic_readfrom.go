package okatype_magic

import (
	"fmt"
	"io"
)

// ReadFrom makes okatype_magic.Type fit the io.ReaderFrom interface.
func (receiver *Type) ReadFrom(r io.Reader) (int64, error) {

	const doesNotMatter = int64(-1)

	if nil == receiver {
		return doesNotMatter, errNilReceiver
	}

	if nil == r {
		return doesNotMatter, errNilReader
	}


	// We try to avoid creating garbage by creating a "bootstrap" backing array
	// for the slice, that is the length that is common for cache lines, at the
	// time this code was initially written.
	var buffer [64]byte
	p := buffer[:]

	if expected := receiver.Len(); len(p) > expected {
		p = p[:expected]
	} else if len(p) < expected {
		p = make([]byte, expected, expected)
	}


	n, err := r.Read(p)
	if nil != err {
		n64 := int64(n)
		if expected, actual := n, int(n64); expected != actual {
			n64 = doesNotMatter
		}
		return n64, fmt.Errorf("Problem reading in magic for Okanero datum: (%T) %q", err, err)
	}

	{
		// We do this (i.e., copy the values over to a local buffer) to prevent the possibility of have the `magic` global being modified.
		//
		// We are OK with using the constant "8" here, because it is a compile time error if the length of `magicCopy` and the `magic` global don't match.
		var magicCopy [8]byte = magic

		for i, expected := range magicCopy {
			if actual := p[i]; expected != actual {
				return doesNotMatter, newBadMagic(magicCopy[:], p)
			}
		}
	}


	var n64 int64 = int64(n)
	if expected, actual := n, int(n64); expected != actual { // Future proof, check for overflow.
		return doesNotMatter, fmt.Errorf("int64 overflow: cannot convert (%T) %#v into and int64.", expected, expected)
	}

	return n64, nil
}
