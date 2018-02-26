/*

With the Okanero protocol, there is a message-oriented layer to the protocol.

Okanero Message

An Okanero Message is represented by okatype_message.Type.

When serialized into a []byte, an Okanero Message can look like:

	[]byte{
		  7,  'o', 'k', 'a', 'n', 'e', 'r', 'o', // <-- magic, as a Pascal-string.
		  1,   0,   0,   0,   0,   0,   0,   0,  // <-- version, as a little-endian 64 bit (8 byte) integer.
		  7,  'm', 'a', 'i', 'n', 'n', 'e', 't', // <-- network, as a Pascal-string.
		255, 108,   1,   0,   0,   0,   0,   0,  // <-- length, as a little-endian 64 bit (8 byte) integer.

		//  ... Okanero Block ...
	}

(Note the nested Okanero Block inside of the Okanero Message.
We will go into more detail about Okanero Blocks shortly.)

You may receive a serialized Okanero Message over a network communication, or stored in (or as part of) a file.

You can think of this as the "native" format of an Okanero Message.

When working with an Okanero Message in Golang code (you probably wouldn't want to work with the serialized Okanero Message
directly, but instead), likely would want to load this kind of data into an okatype_message.Type.

(Also, from the Golang code, likely, you would will not have this serialized Okanero Message data in the form of a []byte,
but instead will have it coming from an io.Reader. So....)

For example:

	var r io.Reader
	
	// ...
	
	var message okatype_message.Type
	
	n64, err := message.ReadFrom(r)

You can then more easily work with the different aspects of a Okanero Message more easily. For example:

	fmt.Printf("Okanero Message version = %d \n", message.Version)


Okanero Block

A Okanero Message wraps an Okanero Block.

(You could (probably safely) conflate an Okanero Block with a "Block" in a "Blockchain".)


Serialized Okanero Datum

A serialized Okanero Datum might look like:

	[]byte{
		  7,  'o', 'k', 'a', 'n', 'e', 'r', 'o', // <-- magic, as a Pascal-string.

		  1,   0,   0,   0,   0,   0,   0,   0,  // <-- version, as a little-endian 64 bit (8 byte) integer.

		  7,  'm', 'a', 'i', 'n', 'n', 'e', 't', // <-- network, as a Pascal-string.

		255, 108,   1,   0,   0,   0,   0,   0,  // <-- length, as a little-endian 64 bit (8 byte) integer.


		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- genesis hash, as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		0x38,0x57,0x30,0x89,0xbf,0x19,0x89,0x8a, // interaction identifier, as a little-endian 512 bit (64 byte) integer.
		0x4b,0x2e,0xa6,0xdc,0x47,0xb0,0x2e,0x23, //
		0x6f,0x65,0x09,0x7f,0xeb,0xf0,0xd9,0x05, //
		0xb4,0x1b,0x83,0xc8,0xed,0xc6,0xab,0x1f, //
		0x0b,0x72,0x05,0x8e,0xd6,0x1e,0x42,0x41, //
		0x07,0x9f,0xff,0xf4,0x38,0xb2,0x50,0xdd, //
		0x33,0x55,0x77,0x62,0xb0,0xc6,0xc7,0xc8, //
		0x05,0xc7,0x4c,0xae,0x1f,0xcb,0x69,0xba, //

		0x47,0xc5,0xcc,0xb1,0xda,0xe1,0x7a,0xd5, // trace identifier, as a little-endian 512 bit (64 byte) integer.
		0x5a,0x58,0x4b,0x65,0xd2,0xf7,0xe5,0x52, //
		0x03,0xc9,0x5b,0x11,0x31,0x41,0xca,0xb2, //
		0xc1,0x3d,0xed,0xfd,0x91,0x1d,0x4e,0x32, //
		0xd1,0xb3,0x87,0x00,0xb4,0x4a,0x2c,0x71, //
		0x35,0xd5,0x58,0x53,0x19,0xee,0xba,0x2a, //
		0xc1,0x56,0x94,0x08,0x34,0x2b,0x93,0x26, //
		0x32,0x1d,0x64,0x27,0xb2,0x06,0x06,0x8f, //

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
