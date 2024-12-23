package guard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const msg = "test message"

func TestEqual(t *testing.T) {
	assert.NotPanics(t, func() { Equal(1, 1, msg) })
	assert.Panics(t, func() { Equal(1, 2, msg) })
}

func TestNotEqual(t *testing.T) {
	assert.NotPanics(t, func() { NotEqual(1, 2, msg) })
	assert.Panics(t, func() { NotEqual(1, 1, msg) })
}

func TestLessThan(t *testing.T) {
	assert.NotPanics(t, func() { LessThan(1, 2, msg) })
	assert.Panics(t, func() { LessThan(2, 1, msg) })
}

func TestLessThanFunc(t *testing.T) {
	assert.NotPanics(t, func() { LessThanFunc(1, 2, func(a, b int) int { return a - b }, msg) })
	assert.Panics(t, func() { LessThanFunc(2, 1, func(a, b int) int { return a - b }, msg) })
}

func TestLessThanEq(t *testing.T) {
	assert.NotPanics(t, func() { LessThanEq(1, 2, msg) })
	assert.NotPanics(t, func() { LessThanEq(2, 2, msg) })
	assert.Panics(t, func() { LessThanEq(2, 1, msg) })
}

func TestLessThanEqFunc(t *testing.T) {
	assert.NotPanics(t, func() { LessThanEqFunc(1, 2, func(a, b int) int { return a - b }, msg) })
	assert.NotPanics(t, func() { LessThanEqFunc(2, 2, func(a, b int) int { return a - b }, msg) })
	assert.Panics(t, func() { LessThanEqFunc(2, 1, func(a, b int) int { return a - b }, msg) })
}

func TestGreaterThan(t *testing.T) {
	assert.NotPanics(t, func() { GreaterThan(2, 1, msg) })
	assert.Panics(t, func() { GreaterThan(1, 2, msg) })
}

func TestGreaterThanFunc(t *testing.T) {
	assert.NotPanics(t, func() { GreaterThanFunc(2, 1, func(a, b int) int { return a - b }, msg) })
	assert.Panics(t, func() { GreaterThanFunc(1, 2, func(a, b int) int { return a - b }, msg) })
}

func TestGreaterThanEq(t *testing.T) {
	assert.NotPanics(t, func() { GreaterThanEq(2, 1, msg) })
	assert.NotPanics(t, func() { GreaterThanEq(2, 2, msg) })
	assert.Panics(t, func() { GreaterThanEq(1, 2, msg) })
}

func TestGreaterThanEqFunc(t *testing.T) {
	assert.NotPanics(t, func() { GreaterThanEqFunc(2, 1, func(a, b int) int { return a - b }, msg) })
	assert.NotPanics(t, func() { GreaterThanEqFunc(2, 2, func(a, b int) int { return a - b }, msg) })
	assert.Panics(t, func() { GreaterThanEqFunc(1, 2, func(a, b int) int { return a - b }, msg) })
}

func TestTrue(t *testing.T) {
	assert.NotPanics(t, func() { True(true, msg) })
	assert.Panics(t, func() { True(false, msg) })
}

func TestFalse(t *testing.T) {
	assert.NotPanics(t, func() { False(false, msg) })
	assert.Panics(t, func() { False(true, msg) })
}

func TestExclusive(t *testing.T) {
	assert.NotPanics(t, func() { Exclusive(true, false, msg) })
	assert.Panics(t, func() { Exclusive(true, true, msg) })
	assert.Panics(t, func() { Exclusive(false, false, msg) })
}

func TestError(t *testing.T) {
	assert.NotPanics(t, func() { Error(assert.AnError, msg) })
	assert.Panics(t, func() { Error(nil, msg) })
}

func TestNoError(t *testing.T) {
	assert.NotPanics(t, func() { NoError(nil, msg) })
	assert.Panics(t, func() { NoError(assert.AnError, msg) })
}

func TestNil(t *testing.T) {
	var nilValue *string
	nonNilValue := "value"
	assert.NotPanics(t, func() { Nil(nil, msg) })
	assert.NotPanics(t, func() { Nil(nilValue, msg) })
	assert.Panics(t, func() { Nil(&nonNilValue, msg) })
}

func TestNotNil(t *testing.T) {
	var nilValue *string
	nonNilValue := "value"
	assert.NotPanics(t, func() { NotNil(&nonNilValue, msg) })
	assert.Panics(t, func() { NotNil(nilValue, msg) })
	assert.Panics(t, func() { NotNil(nil, msg) })
}

func TestZero(t *testing.T) {
	assert.NotPanics(t, func() { Zero(0, msg) })
	assert.Panics(t, func() { Zero(1, msg) })

	assert.NotPanics(t, func() { Zero("", msg) })
	assert.Panics(t, func() { Zero("value", msg) })

	value := "value"
	var nonNilValue = &value
	var nilValue *string = nil
	assert.NotPanics(t, func() { Zero(nilValue, msg) })
	assert.Panics(t, func() { Zero(nonNilValue, msg) })
}

func TestNotZero(t *testing.T) {
	assert.NotPanics(t, func() { NotZero(1, msg) })
	assert.Panics(t, func() { NotZero(0, msg) })

	assert.NotPanics(t, func() { NotZero("value", msg) })
	assert.Panics(t, func() { NotZero("", msg) })

	value := "value"
	var nonNilValue = &value
	var nilValue *string = nil
	assert.NotPanics(t, func() { NotZero(nonNilValue, msg) })
	assert.Panics(t, func() { NotZero(nilValue, msg) })
}
