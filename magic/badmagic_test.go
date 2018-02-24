package okatype_magic

import (
	"io/ioutil"

	"testing"
)

func TestBadMagic(t *testing.T) {

	const expectedMagic = "expected"
	const actualMagic   = "actual"

	{
		err := newBadMagic(nil, nil)
		if nil == err {
			t.Errorf("Did not expect to get nil, but actually got: %v", err)
		}
	}

	{
		err := newBadMagic([]byte(expectedMagic), nil)
		if nil == err {
			t.Errorf("Did not expect to get nil, but actually got: %v", err)
		}
	}

	func() {
		err := newBadMagic(nil, []byte(actualMagic))
		if nil == err {
			t.Errorf("Did not expect to get nil, but actually got: %v", err)
		}

		r := err.Reader()
		if nil == r {
			t.Errorf("Did not expect a nil reader, but actually got: %v", r)
			return
		}

		p, e := ioutil.ReadAll(r)
		if nil != e {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", e, e)
			return
		}

		if expected, actual := actualMagic, string(p); expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}()

	func() {
		err := newBadMagic([]byte(expectedMagic), []byte(actualMagic))
		if nil == err {
			t.Errorf("Did not expect to get nil, but actually got: %v", err)
			return
		}

		r := err.Reader()
		if nil == r {
			t.Errorf("Did not expect a nil reader, but actually got: %v", r)
			return
		}

		p, e := ioutil.ReadAll(r)
		if nil != e {
			t.Errorf("Did not expect an error, but actually got one: (%T) %q", e, e)
			return
		}

		if expected, actual := actualMagic, string(p); expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}()
}
