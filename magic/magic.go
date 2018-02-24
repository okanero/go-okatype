package okatype_magic

// Type represents the magic constant value at the beginning of the serialization of an Okanero datum
// (i.e., a serialization of okatype_datum.Type).
//
// This is used in the okatype_datum.Type struct.
//
// To put some of this into content,....
//
// (And first we'll talk about file formats, even though okatype_magic.Type isn't used in a file format per se.
// It will hopefully make more sense this way.)
//
// On many computer systems (at least at the time of writing this code) one strategy to identify the `type`
// of a file is through its extension.
//
// For example:...
//
// "something.txt" is a "text file" because it ends in the ".txt" extension.
//
// "index.html" is an "Hypertext Markup Language (HTML) file" because it ends in the ".html" extension.
//
// "logo.png" is a "Portable Network Graphics (PNG) (image) file" because it ends in the ".png" extension.
//
// "data.csv" is a "Comma-Separated Values (CSV) file" because it ends in the ".csv" extension.
//
// However, not all computer systems take this approach.
//
// For instance, the strategy used with the HTTP protocol is that the file type is specified using the
// "Content-Type" (HTTP response) header. (Here the file extension tends to get ignored.)
//
// For example:...
//
// "Content-Type: text/html" coming from a URL such as "http://example.com/one/two/three.php"
//
// "Content-Type: "image/png" coming from a URL such as "http://example.com/profile/avatar.php"
//
// "Content-Type: text/html" coming from a URL such as "http://example.com/apple/banana/cherry"
//
// Yet another strategy (and the strategy we are making use of with okatype_magic.Type) is the usage
// of a "magic".
//
// "Magics" are also sometimes called "magic numbers" and "file signatures".
//
// What the magic is is that for some file formats, the first number of bytes in any file of that
// file types have the exact same value, and are used to communicate that that file is of that type.
//
// For example:...
//
// The first 4 bytes of a PDF file are always 37, 80, 68, 70; which (in ASCII) spell out: "%PDF".
//
// The first 6 bytes of an XML file are always 60, 63, 120, 109, 108, 32; which (in ASCII) spell out: "<?xml".
//
// The first 4 bytes of a Java class file are always 202, 254, 202, 190; which aren't meant to spell out in ASCII,
// but in hexadecimal look like: CA FE BA BE.
//
// The okatype_magic.Type use to create a "magic" / "magic numbers" / file signature" for okatype_datum.Type.
type Type struct{}

var (
	// The value of `magic` is the magic value (sequence of bytes) at the beginning of a
	// serialization of an Okanero datum (i.e., a serialization of okatype_datum.Type).
	//
	// (A "serialization" is how you would turn an okatype_datum.Type into a []byte,
	// and then, for example, send that data over the network.)
	//
	// (The use of a "magic" at the beginning of a []byte is a similar technique used by many
	// file formats, where the first number of bytes of it enables you to identify the type of
	// the file it is. Which, for files that follow this convention, allows software to see if
	// it should "bail" or "ignore" it quicker, in O(1) time, by seeing if the "magic" matches
	// what it expects.)
	//
	// The format for the magic is designed to try to be future proof, by allowing to variable
	// length strings. (I.e., the first byte representing the length of the string.)
	//
	// But, at the same time, in the here and now, it is designed for some efficiency, in that
	// it is 8 bytes long, which is the word length on many computer architectures today.
	//
	//         The number of bytes in the string
	//                      │
	//                      │          "okanero"
	//                      │              │
	//                      │   ───────────┴───────────
	//                      │  ╱                       ╲
	//                      ↓ ╱                         ╲
	//              [8]byte{7,'o','k','a','n','e','r','o'}
	magic [8]byte = [8]byte{7,'o','k','a','n','e','r','o'}
)
