package okatype_network

type Type struct {
	loaded bool
	value [8]byte
}

func None() Type {
	return Type{}
}

func Some(value [8]byte) Type {
	return Type{
		value:  value,
		loaded: true,
	}
}
