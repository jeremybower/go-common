package optional

import (
	"bytes"
	"encoding/json"
)

type Map[K comparable, V any] struct {
	Valid bool
	Map   map[K]V
}

func InvalidMap[K comparable, V any]() Map[K, V] {
	return Map[K, V]{Valid: false}
}

func NewMap[K comparable, V any](m map[K]V) Map[K, V] {
	return Map[K, V]{Valid: true, Map: m}
}

func (m *Map[K, V]) Reset() {
	m.Map = nil
	m.Valid = false
}

func (m Map[K, V]) IfValid(fn func(m map[K]V)) {
	if m.Valid {
		fn(m.Map)
	}
}

func (m Map[K, V]) Or(other map[K]V) map[K]V {
	if m.Valid {
		return m.Map
	}

	return other
}

func (m Map[K, V]) OrNil() map[K]V {
	if m.Valid {
		return m.Map
	}

	return nil
}

func (m *Map[K, V]) UnmarshalJSON(b []byte) error {
	if len(b) > 0 {
		if bytes.Equal(b, []byte("null")) {
			return ErrUnexpectedNull
		}

		m.Valid = true
		return json.Unmarshal(b, &m.Map)
	}

	m.Reset()
	return nil
}
