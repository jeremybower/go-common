package guard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert.NotPanics(t, func() { Equal(1, 1) })
	assert.Panics(t, func() { Equal(1, 2) })
}

func TestNotEmpty(t *testing.T) {
	assert.NotPanics(t, func() { NotEmpty(1) })
	assert.Panics(t, func() { NotEmpty("") })
	assert.Panics(t, func() { NotEmpty(nil) })

	makeChan := func(len int) chan int {
		c := make(chan int, len)
		for i := 0; i < len; i++ {
			c <- i
		}
		return c
	}

	assert.Panics(t, func() { NotEmpty(makeChan(0)) })
	assert.NotPanics(t, func() { NotEmpty(makeChan(1)) })

	var nilStr *string = nil
	assert.Panics(t, func() { NotEmpty(nilStr) })

	emptyStr := ""
	assert.Panics(t, func() { NotEmpty(&emptyStr) })

	nonEmptyStr := "hello"
	assert.NotPanics(t, func() { NotEmpty(&nonEmptyStr) })
}

func TestLT(t *testing.T) {
	assert.NotPanics(t, func() { LT(1, 2) })
	assert.Panics(t, func() { LT(2, 1) })
}

func TestLTE(t *testing.T) {
	assert.NotPanics(t, func() { LTE(1, 2) })
	assert.NotPanics(t, func() { LTE(2, 2) })
	assert.Panics(t, func() { LTE(2, 1) })
}

func TestGT(t *testing.T) {
	assert.NotPanics(t, func() { GT(2, 1) })
	assert.Panics(t, func() { GT(1, 2) })
}

func TestGTE(t *testing.T) {
	assert.NotPanics(t, func() { GTE(2, 1) })
	assert.NotPanics(t, func() { GTE(2, 2) })
	assert.Panics(t, func() { GTE(1, 2) })
}

func TestTrue(t *testing.T) {
	assert.NotPanics(t, func() { True(true) })
	assert.Panics(t, func() { True(false) })
}

func TestFalse(t *testing.T) {
	assert.NotPanics(t, func() { False(false) })
	assert.Panics(t, func() { False(true) })
}

func TestXor(t *testing.T) {
	assert.NotPanics(t, func() { Xor(true, false) })
	assert.Panics(t, func() { Xor(true, true) })
	assert.Panics(t, func() { Xor(false, false) })
}

func TestNilOrNotEmpty(t *testing.T) {
	assert.NotPanics(t, func() { NilOrNotEmpty(nil) })
	assert.NotPanics(t, func() { NilOrNotEmpty(1) })
	assert.Panics(t, func() { NilOrNotEmpty(0) })
}

func TestEach(t *testing.T) {
	assert.NotPanics(t, func() { Each([]int{1, 2, 3}, NotNil) })
	assert.Panics(t, func() { Each([]*int{nil, nil, nil}, NotNil) })
}

func TestNotNil(t *testing.T) {
	var nonNilValue = ""
	var nilValue *string = nil
	assert.NotPanics(t, func() { NotNil(nonNilValue) })
	assert.Panics(t, func() { NotNil(nilValue) })
	assert.Panics(t, func() { NotNil(nil) })
}

func TestNotZero(t *testing.T) {
	assert.NotPanics(t, func() { NotZero(1) })
	assert.Panics(t, func() { NotZero(0) })
	assert.Panics(t, func() { NotZero(nil) })
}
