/*
Package okatype_magic provides the okatype_magic.Type type.

With the Okanero protocol, there is a message-oriented layer to it.

The Okanero protocol calls its messages "datum".

(In Golang, we can think of serialized Okanero datum as a []byte.)

(You could probably safely conflate an Okanero datum with a "block" in a "blockchain".)

For safety and future proofing reasons, there is a quick way to identity if something
(such as the payload of a UDP packet) is a Okanero datum or not.

The first 8 bytes of a Okanero datum MUST have the following value:

	[8]byte{7,'o','k','a','n','e','r','o'}

This can be understood as:

	The number of bytes in the string
	             │
	             │          "okanero"
	             │              │
	             │   ───────────┴───────────
	             │  ╱                       ╲
	             ↓ ╱                         ╲
	     [8]byte{7,'o','k','a','n','e','r','o'}

(Some might recognize this style of encoding strings by the following names:
"length-prefixed strings" or "Pascal strings")

The okatype_magic.Type is a way of representing this "magic" (i.e., "magic number")
in a Golang struct. I.e.,

	struct Datum {
		Magic okatype_magic.Type

		// ...
	}

With okatype_magic.Type, the okatype_magic.Type.WriteTo() method allows one to serialize
the the okatype_magic.Type into a io.Writer.

Also with okatype_magic.Type, the okatype_magic.Type.ReadFrom() method allows one to validate
the "magic" (i.e., "magic number") coming from an io.Reader.
*/
package okatype_magic
