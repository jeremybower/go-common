package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Parallel()

	value := []int{1, 2}

	tests := []struct {
		name  string
		opt   Slice[int]
		valid bool
		value []int
	}{
		{"invalid", InvalidSlice[int](), false, nil},
		{"valid", NewSlice(value), true, value},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Slice[int]{Valid: tt.valid, Slice: tt.value}, tt.opt)
			assert.Equal(t, tt.value, tt.opt.OrNil())
		})
	}
}

func TestSliceUnmarshallJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		b           []byte
		initialVal  Slice[string]
		expectedVal Slice[string]
		expectedErr error
	}{
		{"invalid after invalid", []byte{}, InvalidSlice[string](), InvalidSlice[string](), nil},
		{"invalid after value", []byte{}, NewSlice([]string{"test"}), InvalidSlice[string](), nil},
		{"null after invalid", []byte("null"), InvalidSlice[string](), InvalidSlice[string](), ErrUnexpectedNull},
		{"null after value", []byte("null"), NewSlice([]string{"test"}), NewSlice([]string{"test"}), ErrUnexpectedNull},
		{"value after invalid", []byte(`["test"]`), InvalidSlice[string](), NewSlice([]string{"test"}), nil},
		{"value after value", []byte(`["test2"]`), NewSlice([]string{"test"}), NewSlice([]string{"test2"}), nil},
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

func TestSliceIfValid(t *testing.T) {
	t.Parallel()

	Slice[int]{}.IfValid(func(v []int) {
		t.Error("should not be called")
	})

	NewSlice([]int{1, 2}).IfValid(func(v []int) {
		assert.Equal(t, []int{1, 2}, v)
	})
}

func TestSliceOr(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []int{1, 2}, InvalidSlice[int]().Or([]int{1, 2}))
	assert.Equal(t, []int{1, 2}, NewSlice([]int{1, 2}).Or([]int{3, 4}))
}

func TestSliceOrNil(t *testing.T) {
	t.Parallel()

	assert.Nil(t, InvalidSlice[int]().OrNil())
	assert.Equal(t, []int{1, 2}, NewSlice([]int{1, 2}).OrNil())
}
