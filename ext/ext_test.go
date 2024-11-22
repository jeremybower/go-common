package ext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
