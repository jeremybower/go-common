package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Parallel()

	value := map[int]int{1: 2}

	tests := []struct {
		name  string
		opt   Map[int, int]
		valid bool
		value map[int]int
	}{
		{"invalid", InvalidMap[int, int](), false, nil},
		{"valid", NewMap(value), true, value},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Map[int, int]{Valid: tt.valid, Map: tt.value}, tt.opt)
			assert.Equal(t, tt.value, tt.opt.OrNil())
		})
	}
}

func TestMapUnmarshallJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		b           []byte
		initialVal  Map[string, int]
		expectedVal Map[string, int]
		expectedErr error
	}{
		{"invalid after invalid", []byte{}, InvalidMap[string, int](), InvalidMap[string, int](), nil},
		{"invalid after value", []byte{}, NewMap(map[string]int{"test": 1}), InvalidMap[string, int](), nil},
		{"null after invalid", []byte("null"), InvalidMap[string, int](), InvalidMap[string, int](), ErrUnexpectedNull},
		{"null after value", []byte("null"), NewMap(map[string]int{"test": 1}), NewMap(map[string]int{"test": 1}), ErrUnexpectedNull},
		{"value after invalid", []byte(`{"test": 1}`), InvalidMap[string, int](), NewMap(map[string]int{"test": 1}), nil},
		{"value after value", []byte(`{"test": 2}`), NewMap(map[string]int{"test": 1}), NewMap(map[string]int{"test": 2}), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.initialVal
			err := opt.UnmarshalJSON(tt.b)
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedVal, opt)
			}
		})
	}
}

func TestMapIfValid(t *testing.T) {
	t.Parallel()

	Map[int, int]{}.IfValid(func(v map[int]int) {
		t.Error("should not be called")
	})

	NewMap(map[int]int{1: 2}).IfValid(func(v map[int]int) {
		assert.Equal(t, map[int]int{1: 2}, v)
	})
}

func TestMapOr(t *testing.T) {
	t.Parallel()

	assert.Equal(t, map[int]int{1: 2}, InvalidMap[int, int]().Or(map[int]int{1: 2}))
	assert.Equal(t, map[int]int{1: 2}, NewMap(map[int]int{1: 2}).Or(map[int]int{3: 4}))
}

func TestMapOrNil(t *testing.T) {
	t.Parallel()

	assert.Nil(t, InvalidMap[int, int]().OrNil())
	assert.Equal(t, map[int]int{1: 2}, NewMap(map[int]int{1: 2}).OrNil())
}
