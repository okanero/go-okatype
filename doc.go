/*
Package okatype (i.e., Okanero type) provides basic types used for implementing the Okanero protocol, for the Go programming language.


Overview

With the Okanero protocol, there is a message-oriented layer to the protocol.

An Okanero Message is represented, in the Golang code, by: okatype_message.Type.

However, okatype_message.Type makes more sense in the context of three other types.

In the Golang code, there are 4 major types that you should start with, when
dealing with the Okanero protocol's message-oriented layer:

• okatype_message.Type

• okatype_block.Type

• okatype_datum.Type

• okatype_payload.Type

The way that these relate to each other are:

	┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
	┃           Message           ┃
	┃                             ┃
	┃  ┏━━━━━━━━━━━━━━━━━━━━━━━┓  ┃
	┃  ┃         Block         ┃  ┃
	┃  ┃                       ┃  ┃
	┃  ┃  ┏━━━━━━━━━━━━━━━━━┓  ┃  ┃
	┃  ┃  ┃      Datum      ┃  ┃  ┃
	┃  ┃  ┃                 ┃  ┃  ┃
	┃  ┃  ┃  ┏━━━━━━━━━━━┓  ┃  ┃  ┃
	┃  ┃  ┃  ┃  Payload  ┃  ┃  ┃  ┃
	┃  ┃  ┃  ┗━━━━━━━━━━━┛  ┃  ┃  ┃
	┃  ┃  ┗━━━━━━━━━━━━━━━━━┛  ┃  ┃
	┃  ┗━━━━━━━━━━━━━━━━━━━━━━━┛  ┃
	┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛

You can see nesting with okatype_message.Type as:

	package okatype_message

	type Type struct {
		// ...

		Block okatype_block.Type

		// ...
	}

And with okatype_message.Type as:

	package okatype_block

	type Type struct {
		// ...

		Datum okatype_datum.Type

		// ...
	}

And with okatype_datum.Type as:

	package okatype_datum

	type Type struct {
		// ...

		Payload okatype_payload.Type

		// ...
	}

To understand this:

• What is sent over or received from the network is an Okanero Message.

• What is endorsed into a Blockchain is an Okanero Block. This is what an agent acting as an endorser would create. (In other related technologies, you might call "endorsers": "miners" or "stakers".)

• What is created by the author is an Okanero Datum. This is what would-be endorsers would try to turn into an Okanero Block.

• Payload is where the the author would assert any information the author wishes to assert. (Ex: if the payload was an event, then it might have a "name" field (ex: "name"="FOOD_EATEN"), and a "version" field (ex: "verion"="1.0.0"), and possibly some other information (ex: "calories"="240", "sodium"="10g").)


Okanero Message

An Okanero Message is represented by okatype_message.Type.

An Okanero Message is what you would send over or receive from the network.
An Okanero Message is also what you might save to or read from a file.

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

You can then more easily work with the different aspects of an Okanero Message more easily. For example:

	fmt.Printf("Okanero Message version = %d \n", message.Version)


Okanero Block

An Okanero Message wraps an Okanero Block.

(You could (probably safely) conflate an Okanero Block with a "Block" in a "Blockchain".)

When serialized into a []byte, an Okanero Message can look like:

	[]byte{
		// ... Okanero Datum ...

		0x47,0xc5,0xcc,0xb1,0xda,0xe1,0x7a,0xd5, // trace, as a little-endian 512 bit (64 byte) integer.
		0x5a,0x58,0x4b,0x65,0xd2,0xf7,0xe5,0x52, //
		0x03,0xc9,0x5b,0x11,0x31,0x41,0xca,0xb2, //
		0xc1,0x3d,0xed,0xfd,0x91,0x1d,0x4e,0x32, //
		0xd1,0xb3,0x87,0x00,0xb4,0x4a,0x2c,0x71, //
		0x35,0xd5,0x58,0x53,0x19,0xee,0xba,0x2a, //
		0xc1,0x56,0x94,0x08,0x34,0x2b,0x93,0x26, //
		0x32,0x1d,0x64,0x27,0xb2,0x06,0x06,0x8f, //

		0x40,0x47,0x93,0x5A,0x53,0x8b,0xbb,0x7b, // <-- time, as a little-endian 128 it (16 byte) integer.
		0x60,0x98,0x89,0x15,0x2C,0x00,0x00,0x00, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- ensorser, as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- ensorser signature, as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- nonce, as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //
	}

(Note the nested Okanero Datum inside of the Okanero Message.
We will go into more detail about Okanero Datum shortly.)


Serialized Okanero Datum

A serialized Okanero Datum might look like:

	[]byte{
		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- genesis (block hash), as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- prev (block hash), as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		124, 122,   5,   0,   1,   0,   0,   0,  // <-- count, as a little-endian 512 bit (64 byte) integer.
		  0,   0,   0,   0,   0,   0,   0,   0,  //
		  0,   0,   0,   0,   0,   0,   0,   0,  //
		  0,   0,   0,   0,   0,   0,   0,   0,  //
		  0,   0,   0,   0,   0,   0,   0,   0,  //
		  0,   0,   0,   0,   0,   0,   0,   0,  //
		  0,   0,   0,   0,   0,   0,   0,   0,  //
		  0,   0,   0,   0,   0,   0,   0,   0,  //

		0x47,0xc5,0xcc,0xb1,0xda,0xe1,0x7a,0xd5, // interaction, as a little-endian 512 bit (64 byte) integer.
		0x5a,0x58,0x4b,0x65,0xd2,0xf7,0xe5,0x52, //
		0x03,0xc9,0x5b,0x11,0x31,0x41,0xca,0xb2, //
		0xc1,0x3d,0xed,0xfd,0x91,0x1d,0x4e,0x32, //
		0xd1,0xb3,0x87,0x00,0xb4,0x4a,0x2c,0x71, //
		0x35,0xd5,0x58,0x53,0x19,0xee,0xba,0x2a, //
		0xc1,0x56,0x94,0x08,0x34,0x2b,0x93,0x26, //
		0x32,0x1d,0x64,0x27,0xb2,0x06,0x06,0x8f, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- lateral 1 (block hash), as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- lateral 2 (block hash), as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- lateral 3 (block hash), as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //



		// ... Okanero Payload ...



		0x40,0x47,0x93,0x5A,0x53,0x8b,0xbb,0x7b, // <-- time, as a little-endian 128 it (16 byte) integer.
		0x60,0x98,0x89,0x15,0x2C,0x00,0x00,0x00, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- author, as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- author signature, as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //

		0x39,0x2e,0x40,0xed,0x3e,0x6f,0x67,0x13, // <-- nonce, as a little-endian 512 bit (64 byte) integer.
		0xae,0x6d,0xc1,0x83,0xa7,0xed,0x5d,0x9d, //
		0x4f,0x36,0x6d,0x02,0x5b,0xb7,0x8f,0xc7, //
		0x9d,0x5a,0x22,0x71,0x29,0x18,0x56,0x0b, //
		0x80,0x1e,0x90,0xed,0xda,0x03,0x3e,0x3f, //
		0xfb,0x06,0xfb,0x8c,0x85,0x5e,0x02,0x02, //
		0x40,0x41,0xba,0xe1,0xe0,0x63,0x9f,0xb1, //
		0xcf,0x1a,0x98,0x0b,0x28,0xae,0x55,0xb0, //
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

You can then more easily work with the different aspects of an Okanero Datum more easily. For example:

	fmt.Printf("Okanero Datum version = %d \n", datum.Version)

*/
package okatype
