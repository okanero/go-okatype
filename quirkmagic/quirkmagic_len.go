package okatype_magic

// Len returns the length in bytes of the quirk "magic".
func (receiver Type) Len() int {
	return len(quirkmagic)
}
