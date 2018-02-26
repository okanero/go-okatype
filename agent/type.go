package okatype_agent

type Type struct {
	loaded bool
	value [64]byte
}

func None() Type {
	return Type{}
}

func Some(value [64]byte) Type {
	return Type{
		value:  value,
		loaded: true,
	}
}
