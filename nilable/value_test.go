package nilable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	t.Parallel()

	str := "test"

	tests := []struct {
		name  string
		opt   Value[string]
		valid bool
		value *string
	}{
		{"invalid", InvalidValue[string](), false, nil},
		{"valid nil", NewValue[string](nil), true, nil},
		{"valid non-nil", NewValue(&str), true, &str},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Value[string]{Valid: tt.valid, Value: tt.value}, tt.opt)
		})
	}
}

func TestValueUnmarshallJSON(t *testing.T) {
	t.Parallel()

	str1 := "test1"
	str2 := "test2"

	tests := []struct {
		name     string
		b        []byte
		initial  Value[string]
		expected Value[string]
	}{
		{"invalid after invalid", []byte{}, InvalidValue[string](), InvalidValue[string]()},
		{"invalid after value", []byte{}, NewValue(&str1), InvalidValue[string]()},
		{"null after invalid", []byte("null"), InvalidValue[string](), NilValue[string]()},
		{"null after value", []byte("null"), NewValue(&str1), NilValue[string]()},
		{"value after invalid", []byte(`"test1"`), InvalidValue[string](), NewValue(&str1)},
		{"value after value", []byte(`"test2"`), NewValue[string](&str1), NewValue(&str2)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.initial
			err := opt.UnmarshalJSON(tt.b)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, opt)
		})
	}
}

func TestValueIfValid(t *testing.T) {
	t.Parallel()

	str := "test"

	Value[string]{}.IfValid(func(v *string) {
		t.Error("should not be called")
	})

	NewValue(&str).IfValid(func(v *string) {
		assert.Equal(t, &str, v)
	})
}

func TestValueOr(t *testing.T) {
	t.Parallel()

	str := "test"

	assert.Nil(t, InvalidValue[string]().Or(nil))
	assert.Equal(t, &str, InvalidValue[string]().Or(&str))
	assert.Equal(t, &str, NewValue(&str).Or(nil))
}

func TestValueOrNil(t *testing.T) {
	t.Parallel()

	str := "test"

	assert.Nil(t, InvalidValue[string]().OrNil())
	assert.Equal(t, &str, NewValue(&str).OrNil())
}
