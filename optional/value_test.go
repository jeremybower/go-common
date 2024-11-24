package optional

import (
	"testing"

	. "github.com/jeremybower/go-common/ext"
	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	t.Parallel()

	value := 42

	tests := []struct {
		name       string
		opt        Value[int]
		valid      bool
		value      int
		valueOrNil *int
	}{
		{"invalid", InvalidValue[int](), false, 0, nil},
		{"valid", NewValue(value), true, value, &value},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Value[int]{Valid: tt.valid, Value: tt.value}, tt.opt)
			assert.Equal(t, tt.valueOrNil, tt.opt.OrNil())
		})
	}
}

func TestValueUnmarshallJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		b           []byte
		initialVal  Value[string]
		expectedVal Value[string]
		expectedErr error
	}{
		{"invalid after invalid", []byte{}, InvalidValue[string](), InvalidValue[string](), nil},
		{"invalid after value", []byte{}, NewValue("test"), InvalidValue[string](), nil},
		{"null after invalid", []byte("null"), InvalidValue[string](), InvalidValue[string](), ErrUnexpectedNull},
		{"null after value", []byte("null"), NewValue("test"), NewValue("test"), ErrUnexpectedNull},
		{"value after invalid", []byte(`"test"`), InvalidValue[string](), NewValue("test"), nil},
		{"value after value", []byte(`"test2"`), NewValue[string]("test1"), NewValue("test2"), nil},
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

func TestValueIfValid(t *testing.T) {
	t.Parallel()

	Value[int]{}.IfValid(func(v int) {
		t.Error("should not be called")
	})

	NewValue(42).IfValid(func(v int) {
		assert.Equal(t, 42, v)
	})
}

func TestValueOr(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 42, InvalidValue[int]().Or(42))
	assert.Equal(t, 42, NewValue(42).Or(43))
}

func TestValueOrNil(t *testing.T) {
	t.Parallel()

	assert.Nil(t, InvalidValue[int]().OrNil())
	assert.Equal(t, Ptr(42), NewValue(42).OrNil())
}
