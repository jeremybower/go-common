package optional

import (
	"bytes"
	"encoding/json"
)

type Value[T any] struct {
	Valid bool
	Value T
}

func InvalidValue[T any]() Value[T] {
	return Value[T]{Valid: false}
}

func NewValue[T any](v T) Value[T] {
	return Value[T]{Valid: true, Value: v}
}

func (v *Value[T]) Reset() {
	var zero T
	v.Value = zero
	v.Valid = false
}

func (v Value[T]) IfValid(fn func(v T)) {
	if v.Valid {
		fn(v.Value)
	}
}

func (v Value[T]) Or(other T) T {
	if v.Valid {
		return v.Value
	}

	return other
}

func (v Value[T]) OrNil() *T {
	if v.Valid {
		return &v.Value
	}

	return nil
}

func (v *Value[T]) UnmarshalJSON(b []byte) error {
	if len(b) > 0 {
		if bytes.Equal(b, []byte("null")) {
			return ErrUnexpectedNull
		}

		v.Valid = true
		return json.Unmarshal(b, &v.Value)
	}

	v.Reset()
	return nil
}
