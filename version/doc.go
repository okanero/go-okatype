/*
Package okatype_version provides the okatype_version.Type and okatype_version.NullableType types.

With the Okanero protocol, there is a message-oriented layer to it.

The Okanero protocol calls its messages "datum".

(In Golang, we can think of serialized Okanero datum as a []byte.)

(You could probably safely conflate an Okanero datum with a "block" in a "blockchain".)

For safety and future proofing reasons, there is a mechanism to allow different versions of an Okanero datum.

The second 8 bytes (after the magic) of a Okanero datum is the Okanero datum version.

Those 8 bytes are interpreted as a little-endian 64 bit integer.

This can be understood as:

	[8]byte{203,255,0,0,0,0,0,0}

... begin interpreted as:

	(203 × 256^0)  +  (255 × 256^1) =

	(203 ×   1)    +  (255 × 256)   =

	 203           +  65280         =

	65513

So, continuing our example, if we include the magic also, the first 16 bytes in that example would be:

	[16]byte{
		7,'o','k','a','n','e','r','o' // <-- magic, as a Pascal-string
		203,255,0,0,0,0,0,0,          // <-- version, as a little-endian 64 bit integer.
	}

MOST LIKELY YOU WOULD NOT CREATE YOUR OWN okatype_version.Type, BUT INSTEAD USE IT FROM AN okatype_datum.Type.
*/
package okatype_version
