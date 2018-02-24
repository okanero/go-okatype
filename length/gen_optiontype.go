package okatype_length

/*
 * CODE GENERATED AUTOMATICALLY WITH https://github.com/reiver/gogen-optiontype
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var (
	errNilReceiver  = errors.New("Nil Receiver")
	errNone         = errors.New("okatype_length.None()")
	errNoneNullable = errors.New("okatype_length.NoneNullable()")
	errNull         = errors.New("okatype_length.Null()")
)

func (receiver NullableType) Int64() (int64, error) {
	if NoneNullable() == receiver {
		return 0, errNoneNullable
	}
	if Null() == receiver {
		return 0, errNull
	}

	return receiver.value, nil
}

func (receiver *NullableType) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == src {
		*receiver = Null()
		return nil
	}

	switch t := src.(type) {
	case NullableType:
		*receiver = t
		return nil
	case Type:
		switch t {
		case None():
			*receiver = NoneNullable()
		default:
			datum, err := t.Int64()
			if nil != err {
				return fmt.Errorf("Problem unwrapping %T: (%T) %v", t, err, err)
			}
			*receiver = SomeNullable(datum)
		}
		return nil
	case int64:
		*receiver = SomeNullable(t)
		return nil
	case string:
		i64, err := strconv.ParseInt(t, 10, 64)
		if nil != err {
			return err
		}
		*receiver = SomeNullable(i64)
		return nil
	case []byte:
		s := string(t)
		i64, err := strconv.ParseInt(s, 10, 64)
		if nil != err {
			return err
		}
		*receiver = SomeNullable(i64)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}

func (receiver NullableType) String() string {
	if NoneNullable() == receiver {
		return "okatype_length.NoneNullable()"
	}
	if Null() == receiver {
		return "okatype_length.Null()"
	}

	return fmt.Sprintf("okatype_length.SomeNullable(%d)", receiver.value)
}

type NullableType struct {
	loaded bool
	null   bool
	value  int64
}

func (receiver NullableType) MarshalJSON() ([]byte, error) {
	if NoneNullable() == receiver {
		return nil, errNoneNullable
	}
	if Null() == receiver {
		return json.Marshal(nil)
	}

	return json.Marshal(receiver.value)
}

func (receiver NullableType) WhenNone(fn func()) {
	if NoneNullable() == receiver {
		fn()
	}
}

func (receiver NullableType) WhenNull(fn func()) {
	if Null() == receiver {
		fn()
	}
}

func (receiver NullableType) WhenSome(fn func(int64)) {
	if NoneNullable() != receiver && Null() != receiver {
		fn(receiver.value)
	}
}

func (receiver *NullableType) UnmarshalJSON(b []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 0 == bytes.Compare(b, []byte{'n','u','l','l'}) {
		*receiver = Null()
		return nil
	}

	var target int64

	if err := json.Unmarshal(b, &target); nil != err {
		return err
	}

	*receiver = SomeNullable(target)

	return nil
}

func (receiver NullableType) Value() (driver.Value, error) {
	if NoneNullable() == receiver {
		return nil, errNoneNullable
	}
	if Null() == receiver {
		return nil, nil
	}

	return receiver.value, nil
}

func NoneNullable() NullableType {
	return NullableType{}
}

func SomeNullable(value int64) NullableType {
	return NullableType{
		value:  value,
		loaded: true,
	}
}

func Null() NullableType {
	return NullableType{
		null:   true,
		loaded: true,
	}
}

func (receiver Type) Int64() (int64, error) {
	if None() == receiver {
		return 0, errNone
	}

	return receiver.value, nil
}

func (receiver *Type) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch t := src.(type) {
	case Type:
		*receiver = t
		return nil
	case int64:
		*receiver = Some(t)
		return nil
	case string:
		i64, err := strconv.ParseInt(t, 10, 64)
		if nil != err {
			return err
		}
		*receiver = Some(i64)
		return nil
	case []byte:
		s := string(t)
		i64, err := strconv.ParseInt(s, 10, 64)
		if nil != err {
			return err
		}
		*receiver = Some(i64)
		return nil
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}

func (receiver Type) String() string {
	if None() == receiver {
		return "okatype_length.None()"
	}

	return fmt.Sprintf("okatype_length.Some(%d)", receiver.value)
}

type Type struct {
	loaded bool
	value  int64
}

func (receiver Type) MarshalJSON() ([]byte, error) {
	if None() == receiver {
		return nil, errNone
	}

	return json.Marshal(receiver.value)
}

func (receiver *Type) UnmarshalJSON(b []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if 0 == bytes.Compare(b, []byte{'n','u','l','l'}) {
		return fmt.Errorf("Cannot unmarshal %q into %T.", string(b), *receiver)
	}

	var target int64

	if err := json.Unmarshal(b, &target); nil != err {
		return err
	}

	*receiver = Some(target)

	return nil
}

func (receiver Type) WhenNone(fn func()) {
	if None() == receiver {
		fn()
	}
}

func (receiver Type) WhenSome(fn func(int64)) {
	if None() != receiver {
		fn(receiver.value)
	}
}

func (receiver Type) Value() (driver.Value, error) {
	if None() == receiver {
		return nil, errNone
	}

	return receiver.value, nil
}

func None() Type {
	return Type{}
}

func Some(value int64) Type {
	return Type{
		value:  value,
		loaded: true,
	}
}
