package okatype_pascalstring

type Type struct {
	loaded  bool
	value string
}

func None() Type {
	return Type{}
}

func Some(value string) Type {
	return Type{
		value:  value,
		loaded: true,
	}
}
