package okatype_quirkmagic

import (
	"io/ioutil"

	"testing"
)

func TestBadQuirkMagic(t *testing.T) {

	const expectedQuirkMagic = "expected"
	const actualQuirkMagic   = "actual"

	{
		err := newBadQuirkMagic(nil, nil)
		if nil == err {
			t.Errorf("Did not expect to get nil, but actually got: %v", err)
		}
	}

	{
		err := newBadQuirkMagic([]byte(expectedQuirkMagic), nil)
		if nil == err {
			t.Errorf("Did not expect to get nil, but actually got: %v", err)
		}
	}

	func() {
		err := newBadQuirkMagic(nil, []byte(actualQuirkMagic))
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

		if expected, actual := actualQuirkMagic, string(p); expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}()

	func() {
		err := newBadQuirkMagic([]byte(expectedQuirkMagic), []byte(actualQuirkMagic))
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

		if expected, actual := actualQuirkMagic, string(p); expected != actual {
			t.Errorf("Expected %q, but actually got %q.", expected, actual)
			return
		}
	}()
}
