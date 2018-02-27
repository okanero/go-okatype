package okatype_pascalstring

import (
	"io"
)

func (receiver Type) WriteTo(w io.Writer) (n int64, err error) {
	const doesNotMatter = -1

	if nil == w {
		return doesNotMatter, errNilWriter
	}

	var n64 int64 = 0

	{
		length := len(receiver.value)

		if maxPascalStringLength < length {
			return doesNotMatter, errTooLong
		}

		var buffer [1]byte

		buffer[0] = byte(length)
		p := buffer[:]

		n, err := w.Write(p)
		if nil != err {
			return doesNotMatter, err
		}

		n64 += int64(n)
	}

	{
		n, err := io.WriteString(w, receiver.value)
		if nil != err {
			return doesNotMatter, err
		}

		n64 += int64(n)
	}


	return n64, nil
}
