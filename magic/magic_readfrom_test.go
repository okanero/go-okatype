package okatype_magic

import (
	"bytes"
	"io"
	"io/ioutil"

	"testing"
)

func TestMagicReaderFrom(t *testing.T) {
	var datum Type

	var x io.ReaderFrom = &datum // THIS IS WHAT ACTUALLY MATTERS.

	if nil == x {
		t.Errorf("This should never happen.")
	}
}

func TestMagicReadFromNilReader(t *testing.T) {

	var x Type

	_, err := x.ReadFrom(nil)
	if nil == err {
		t.Errorf("Expected an error, but did not actually get one: %q", err)
		return
	}
	if expected, actual := errNilReader, err; expected != actual {
		t.Errorf("Received wrong error, expected (%T) %q, but actually got (%T) %q", expected, expected, actual, actual)
		return
	}
}

func TestMagicReadFromBadMagic(t *testing.T) {

	tests := []struct{
		Magic []byte
	}{
		{
			Magic: []byte{ 7,'t','h','i','s','i','s','a'},
		},

		{
			Magic: []byte{11,'h','e','l','l','o',' ','w','o','r','l','d'},
		},

		{
			Magic: []byte{
				' ','!','"','#','$','%','&','\'','(',')','*','+',',','-','.','/',
				'0','1','2','3','4','5','6','7','8','9',
				':',';','<','=','>','?','@',
				'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','U','Z',
				'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z',
			},
		},

		{
			Magic: []byte{'<','?','x',',','l',' ','v','e'},
		},

		{
			Magic: []byte{'<','!','D','O','C','T','Y',','},
		},

		{
			Magic: []byte{7,'O','K','A','N','E','R','O'},
		},
		{
			Magic: []byte{7,'O','k','a','n','e','r','o'},
		},
		{
			Magic: []byte{7,'O','k','A','n','E','r','O'},
		},
		{
			Magic: []byte{7,'o','K','a','N','e','R','o'},
		},

		{
			Magic: []byte{7,'0','k','a','n','e','r','o'},
		},
		{
			Magic: []byte{7,'o','k','a','n','e','r','0'},
		},
		{
			Magic: []byte{7,'0','k','a','n','e','r','0'},
		},
	}
	for i := byte(0); i <= byte(20); i++ {
		if i == magic[0] {
			continue
		}

		test := struct {
			Magic []byte
		}{
			Magic: []byte{i,'o','k','a','n','e','r','o'},
		}

		tests = append(tests, test)
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		if expected, actual := 0, buffer.Len(); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			return
		}

		{
			n, err := buffer.Write(test.Magic)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
				continue
			}
			if expected, actual := len(test.Magic), n; expected != actual {
				t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
				continue
			}
		}

		{
			var magic Type

			_, err :=  magic.ReadFrom(&buffer)
			if nil == err {
				t.Errorf("For test #%d, expected an error, but did not actually get one: %v", testNumber, err)
				continue
			}
			switch complainer := err.(type) {
			default:
				t.Errorf("For test #%d, expected \"bad magic\" error, but actually got: (%T) %v", testNumber, err, err)
			case BadMagic:
				r := complainer.Reader()
				if nil == r {
					t.Errorf("For test #%d, did not expect a nil reader, but actually got one: %v", testNumber, r)
					continue
				}
				if expected, actual := magic.Len(), r.Len(); expected != actual {
					t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
					continue
				}
				{
					p, err := ioutil.ReadAll(r)
					if nil != err {
						t.Errorf("For test #%d, did not expected an error, but actually got one: (%T) %v", testNumber, err, err)
						continue
					}
					if expected, actual := string(test.Magic[:len(p)]), string(p); expected != actual {
						t.Errorf("For test #%d, expected %#v, but actually got %#v.", testNumber, expected, actual)
						continue

					}
				}
			}
		}
	}
}
