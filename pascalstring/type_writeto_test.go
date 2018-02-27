package okatype_pascalstring

import (
	"bytes"
	"io"

	"testing"
)

func TestTypeWriterTo(t *testing.T) {

	var x io.WriterTo = Type{} // THIS IS WHAT ACTUALLY MATTERS!

	if  nil == x{
		t.Errorf("This should never happen.")
	}
}

func TestTypeWriteTo(t *testing.T) {

	tests := []struct{
		String     string
		Expected []byte
	}{
		{
			String:   "",
			Expected: []byte{0},
		},



		{
			String:   "apple",
			Expected: []byte{5, 'a','p','p','l','e'},
		},
		{
			String:   "banana",
			Expected: []byte{6, 'b','a','n','a','n','a'},
		},
		{
			String:   "cherry",
			Expected: []byte{6, 'c','h','e','r','r','y'},
		},



		{
			String:   "Hello world!",
			Expected: []byte{12, 'H','e','l','l','o',' ','w','o','r','l','d','!'},
		},



		{
			String:   "I am Darius, the great king, king of kings, the king of Persia, the king of countries, the son of Hystaspes, the grandson of Arsames, the Achaemenid.",
			Expected: []byte{149, 'I',' ','a','m',' ','D','a','r','i','u','s',',',' ','t','h','e',' ','g','r','e','a','t',' ','k','i','n','g',',',' ','k','i','n','g',' ','o','f',' ','k','i','n','g','s',',',' ','t','h','e',' ','k','i','n','g',' ','o','f',' ','P','e','r','s','i','a',',',' ','t','h','e',' ','k','i','n','g',' ','o','f',' ','c','o','u','n','t','r','i','e','s',',',' ','t','h','e',' ','s','o','n',' ','o','f',' ','H','y','s','t','a','s','p','e','s',',',' ','t','h','e',' ','g','r','a','n','d','s','o','n',' ','o','f',' ','A','r','s','a','m','e','s',',',' ','t','h','e',' ','A','c','h','a','e','m','e','n','i','d','.'},
		},



		{
			String:   "[0]4567890[1]4567890[2]4567890[3]4567890[4]4567890[5]4567890[6]4567890[7]4567890[8]4567890[9]4567890[10]567890[11]567890[12]567890[13]567890[14]567890[15]567890[16]567890[17]567890[18]567890[19]567890[20]567890[21]567890[22]567890[23]567890[24]567890[25]5",
			Expected: []byte{255, '[','0',']','4','5','6','7','8','9','0','[','1',']','4','5','6','7','8','9','0','[','2',']','4','5','6','7','8','9','0','[','3',']','4','5','6','7','8','9','0','[','4',']','4','5','6','7','8','9','0','[','5',']','4','5','6','7','8','9','0','[','6',']','4','5','6','7','8','9','0','[','7',']','4','5','6','7','8','9','0','[','8',']','4','5','6','7','8','9','0','[','9',']','4','5','6','7','8','9','0','[','1','0',']','5','6','7','8','9','0','[','1','1',']','5','6','7','8','9','0','[','1','2',']','5','6','7','8','9','0','[','1','3',']','5','6','7','8','9','0','[','1','4',']','5','6','7','8','9','0','[','1','5',']','5','6','7','8','9','0','[','1','6',']','5','6','7','8','9','0','[','1','7',']','5','6','7','8','9','0','[','1','8',']','5','6','7','8','9','0','[','1','9',']','5','6','7','8','9','0','[','2','0',']','5','6','7','8','9','0','[','2','1',']','5','6','7','8','9','0','[','2','2',']','5','6','7','8','9','0','[','2','3',']','5','6','7','8','9','0','[','2','4',']','5','6','7','8','9','0','[','2','5',']','5'},
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var s Type = Some(test.String)

		n64, err := s.WriteTo(&buffer)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}
		if expected, actual := int64(1+len(test.String)), n64; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		if expected, actual := test.Expected, buffer.Bytes(); !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, expected (%d, %q), but actually got (%d, %q)", testNumber, expected[0], expected[1:], actual[0], actual[1:])
			continue
		}
	}
}

func TestTypeWriteToError(t *testing.T) {

	tests := []struct{
		String string
	}{
		{
			String: "[0]4567890[1]4567890[2]4567890[3]4567890[4]4567890[5]4567890[6]4567890[7]4567890[8]4567890[9]4567890[10]567890[11]567890[12]567890[13]567890[14]567890[15]567890[16]567890[17]567890[18]567890[19]567890[20]567890[21]567890[22]567890[23]567890[24]567890[25]56",
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		var s Type = Some(test.String)

		_, err := s.WriteTo(&buffer)
		if nil == err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}
		if expected, actual := errTooLong, err; expected != actual {
			t.Errorf("For test #%d, expected (%T) %q, but actually got (%T) %q", testNumber, expected, expected, actual, actual)
			continue
		}
	}
}
