/*
Package okatype (i.e., Okanero type) provides basic types used for implementing the Okanero protocol.

With the Okanero protocol, there is a message-oriented layer to it.

The Okanero protocol calls its messages "Datum".

(In Golang, we can think of serialized Okanero Datum as a []byte.)

(You could probably safely conflate an Okanero Datum with a "Block" in a "Blockchain".)

Serialized Okanero Datum

A serialized Okanero Datum might look like:

	[]byte{
		7, 'o','k','a','n','e','r','o', // <-- magic, as a Pascal-string
		1,  0,  0,  0,  0,  0,  0,  0,  // <-- version, as a little-endian 64 bit (8 byte) integer.

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

	fmt.Printf("Okanero Datum version = %d", datum.Version)

*/
package okatype
