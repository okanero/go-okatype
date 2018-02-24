/*
Package okatype_length (i.e., Okanero type length) provides the okatype_length.Type and okatype_length.NullableType types.

With the Okanero protocol, there is a message-oriented layer to it.

The Okanero protocol calls its messages "datum".

(In Golang, we can think of serialized Okanero datum as a []byte.)

(You could probably safely conflate an Okanero datum with a "block" in a "blockchain".)

A Okanero datum is variable length, thus there is a mechanism to specify the length of a Okanero datum.

The third 8 bytes (after the magic and the version) of a Okanero datum is the Okanero datum length.

Those 8 bytes are interpreted as a little-endian 64 bit integer.

This can be understood as:

	[8]byte{203,255,0,0,0,0,0,0}

... begin interpreted as:

	(203 × 256^0)  +  (255 × 256^1) =

	(203 ×   1)    +  (255 × 256)   =

	 203           +  65280         =

	65513

So, continuing our example, if we include the magic and version also, the first 32 bytes in that example would be:

	[32]byte{
		7,'o','k','a','n','e','r','o' // <-- magic, as a Pascal-string
		1,0,0,0,0,0,0,0,              // <-- version, as a little-endian 64 bit integer.
		203,255,0,0,0,0,0,0,          // <-- length, as a little-endian 64 bit integer.
	}

The okatype_length.Type is a way of representing this "length" in a Golang struct. I.e.,

	struct Datum {
		Magic   okatype_magic.Type
		Version okatype_version.Type
		Length  okatype_length.Type

		// ...
	}

MOST LIKELY YOU WOULD NOT CREATE YOUR OWN okatype_length.Type, BUT INSTEAD USE IT FROM AN okatype_length.Type.
*/
package okatype_length
