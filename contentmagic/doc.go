/*
Package okatype_contentmagic (i.e., Okanero type content magic) provides okatype_contentmagic.Type.

The first 8 bytes of an Okanero Content MUST have the following bytes:

	[8]byte{7,'c','o','n','t','e','n','t'}

This can be understood as:

	The number of bytes in the string
	             │
	             │          "content"
	             │              │
	             │   ───────────┴───────────
	             │  ╱                       ╲
	             ↓ ╱                         ╲
	     [8]byte{7,'c','o','n','t','e','n','t'}

(Some might recognize this style of encoding strings by the following names:
"length-prefixed strings" or "Pascal strings")

The okatype_contentmagic.Type is a way of representing this "magic" (i.e., "magic number")
in a Golang struct. I.e.,

	struct Content {
		Magic okatype_contentmagic.Type

		// ...
	}

With okatype_contentmagic.Type, the okatype_contentmagic.Type.WriteTo() method allows one to serialize
the the okatype_contentmagic.Type into a io.Writer.

Also with okatype_contentmagic.Type, the okatype_contentmagic.Type.ReadFrom() method allows one to validate
the "magic" (i.e., "magic number") coming from an io.Reader.

MOST LIKELY YOU WOULD NOT CREATE YOUR OWN okatype_contentmagic.Type, BUT INSTEAD USE IT FROM AN okatype_content.Type.
*/
package okatype_contentmagic
