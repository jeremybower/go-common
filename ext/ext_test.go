package ext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	assert.Equal(t, 1, If(true, 1, 2))
	assert.Equal(t, 2, If(false, 1, 2))
}

func TestPtr(t *testing.T) {
	str := "hello"
	assert.Equal(t, &str, Ptr(str))

	num := 123
	assert.Equal(t, &num, Ptr(num))
}

func TestPtrNil(t *testing.T) {
	var str *string
	assert.Equal(t, str, NilPtr[string]())
}
