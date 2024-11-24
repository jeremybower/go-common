package optional

import (
	"bytes"
	"encoding/json"
)

type Slice[T any] struct {
	Valid bool
	Slice []T
}

func InvalidSlice[T any]() Slice[T] {
	return Slice[T]{Valid: false}
}

func NewSlice[T any](s []T) Slice[T] {
	return Slice[T]{Valid: true, Slice: s}
}

func (s *Slice[T]) Reset() {
	s.Slice = nil
	s.Valid = false
}

func (s Slice[T]) IfValid(fn func(s []T)) {
	if s.Valid {
		fn(s.Slice)
	}
}

func (s Slice[T]) Or(other []T) []T {
	if s.Valid {
		return s.Slice
	}

	return other
}

func (s Slice[T]) OrNil() []T {
	if s.Valid {
		return s.Slice
	}

	return nil
}

func (s *Slice[T]) UnmarshalJSON(b []byte) error {
	if len(b) > 0 {
		if bytes.Equal(b, []byte("null")) {
			return ErrUnexpectedNull
		}

		s.Valid = true
		return json.Unmarshal(b, &s.Slice)
	}

	s.Reset()
	return nil
}
