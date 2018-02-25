/*
Package okatype (i.e., Okanero type) provides basic types used for implementing the Okanero protocol.

With the Okanero protocol, there is a message-oriented layer to it.

The Okanero protocol calls its messages "Datum".

(In Golang, we can think of serialized Okanero Datum as a []byte.)

(You could probably safely conflate an Okanero Datum with a "Block" in a "Blockchain".)

Serialized Okanero Datum

A serialized Okanero Datum might look like:

	[]byte{
		   7, 'o', 'k', 'a', 'n', 'e', 'r', 'o', // <-- magic, as a Pascal-string
		   1,  0,   0,   0,   0,   0,   0,   0,  // <-- version, as a little-endian 64 bit (8 byte) integer.

		0x38,0x57,0x30,0x89,0xbf,0x19,0x89,0x8a, // interaction identifier, as a 512 bit (64 byte) integer.
		0x4b,0x2e,0xa6,0xdc,0x47,0xb0,0x2e,0x23, //
		0x6f,0x65,0x09,0x7f,0xeb,0xf0,0xd9,0x05, //
		0xb4,0x1b,0x83,0xc8,0xed,0xc6,0xab,0x1f, //
		0x0b,0x72,0x05,0x8e,0xd6,0x1e,0x42,0x41, //
		0x07,0x9f,0xff,0xf4,0x38,0xb2,0x50,0xdd, //
		0x33,0x55,0x77,0x62,0xb0,0xc6,0xc7,0xc8, //
		0x05,0xc7,0x4c,0xae,0x1f,0xcb,0x69,0xba, //

		// ...
	}

You may receive a serialized Okanero Datum over a network communication, or stored in (or as part of) a file.

You can think of this as the "native" format of an Okanero Datum.

Okanero Datum Struct

When working with an Okanero Datum in Golang code (you probably wouldn't want to work with the serialized Okanero Datum
directly, but instead), likely would want to load this kind of data into an okatype_datum.Type.

(Also, from the Golang code, likely, you would will not have this serialized Okanero Datum data in the form of a []byte,
but instead will have it coming from an io.Reader. So....)

For example:

	var r io.Reader
	
	// ...
	
	var datum okatype_datum.Type
	
	n64, err := datum.ReadFrom(r)

You can then more easily work with the different aspects of a Okanero Datum more easily. For example:

	fmt.Printf("Okanero Datum version = %d \n", datum.Version)

*/
package okatype
