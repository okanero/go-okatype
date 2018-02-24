package okatype_magic

import (
	"bytes"
	"io"

	"testing"
)

func TestMagicWriteToWriter(t *testing.T) {

	var datum Type

	var x io.WriterTo = datum // THIS IS WHAT ACTUALLY MATTERS.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestMagicWriteToNilWriter(t *testing.T) {

	var x Type

	_, err := x.WriteTo(nil)
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %q", err)
		return
	}
	if expected, actual := errNilWriter, err; expected != actual {
		t.Errorf("Received wrong error, expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}
}

func TestMagicWriteTo(t *testing.T) {

	var magic Type

	var buffer bytes.Buffer

	if expected, actual := 0, buffer.Len(); expected != actual {
		t.Errorf("Expected %d, but actually got %d.", expected, actual)
		return
	}

	n64, err := magic.WriteTo(&buffer)
	if nil != err {
		t.Errorf("Did not expect an error, but did not actually get one: (%T) %q", err, err)
		return
	}

	{
		// We hard code the expected length here, because theoretically `Write()` could modify
		// the memory space containing the magic, so len(magic) could be wrong.
		const expected = 8

		if actual := buffer.Len(); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}

		if actual := int(n64); expected != actual {
			t.Errorf("Expected %d, but actually got %d.", expected, actual)
			return
		}
	}

	{
		// We hard code the expected length here, because theoretically `Write()` could modify
		// the memory space containing the magic, so the data could be wrong.
		const expected0 byte = 7
		const expected1 byte = 'o'
		const expected2 byte = 'k'
		const expected3 byte = 'a'
		const expected4 byte = 'n'
		const expected5 byte = 'e'
		const expected6 byte = 'r'
		const expected7 byte = 'o'

		p := buffer.Bytes()

		if expected, actual := expected0, p[0]; expected != actual {
			t.Errorf("Bad magic at byte 0, expected %d, but actually got %d.", expected, expected, actual, actual)
			return
		}
		if expected, actual := expected1, p[1]; expected != actual {
			t.Errorf("Bad magic at byte 1, expected %q (%d), but actually got %q (%d).", expected, expected, actual, actual)
			return
		}
		if expected, actual := expected2, p[2]; expected != actual {
			t.Errorf("Bad magic at byte 2, expected %q (%d), but actually got %q (%d).", expected, expected, actual, actual)
			return
		}
		if expected, actual := expected3, p[3]; expected != actual {
			t.Errorf("Bad magic at byte 3, expected %q (%d), but actually got %q (%d).", expected, expected, actual, actual)
			return
		}
		if expected, actual := expected4, p[4]; expected != actual {
			t.Errorf("Bad magic at byte 4, expected %q (%d), but actually got %q (%d).", expected, expected, actual, actual)
			return
		}
		if expected, actual := expected5, p[5]; expected != actual {
			t.Errorf("Bad magic at byte 5, expected %q (%d), but actually got %q (%d).", expected, expected, actual, actual)
			return
		}
		if expected, actual := expected6, p[6]; expected != actual {
			t.Errorf("Bad magic at byte 6, expected %q (%d), but actually got %q (%d).", expected, expected, actual, actual)
			return
		}
		if expected, actual := expected7, p[7]; expected != actual {
			t.Errorf("Bad magic at byte 7, expected %q (%d), but actually got %q (%d).", expected, expected, actual, actual)
			return
		}
	}
}
