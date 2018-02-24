/*
Package okatype (i.e., Okanero type) provides basic types used for implementing the Okanero protocol.

With the Okanero protocol, there is a message-oriented layer to it.

The Okanero protocol calls its messages "datum".

(In Golang, we can think of serialized Okanero datum as a []byte.)

(You could probably safely conflate an Okanero datum with a "block" in a "blockchain".)

Serialized Datum

A serialized Okranero datum might look like:

	[]byte{
		7, 'o','k','a','n','e','r','o', // <-- magic, as a Pascal-string
		1,  0,  0,  0,  0,  0,  0,  0,  // <-- version, as a little-endian 64 bit integer.

		// ...
	}

*/
package okatype
