package okatype_contentmagic

// Len returns the length in bytes of the "magic".
func (receiver Type) Len() int {
	return len(magic)
}
