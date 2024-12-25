package guard

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

// ----------------------------------------------------------------------------
// Boolean
// ----------------------------------------------------------------------------

func True(value bool, msg string) {
	if !value {
		panic(msg)
	}
}

func False(value bool, msg string) {
	if value {
		panic(msg)
	}
}

func Exclusive(a, b bool, msg string) {
	if !((a || b) && !(a && b)) {
		panic(msg)
	}
}

// ----------------------------------------------------------------------------
// Equality
// ----------------------------------------------------------------------------

func Equal[T comparable](a, b T, msg string) {
	if a != b {
		panic(msg)
	}
}

func NotEqual[T comparable](a, b T, msg string) {
	if a == b {
		panic(msg)
	}
}

// ----------------------------------------------------------------------------
// Error
// ----------------------------------------------------------------------------

func Error(err error, msg string) {
	if err == nil {
		panic(msg)
	}
}

func NoError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

// ----------------------------------------------------------------------------
// Nil
// ----------------------------------------------------------------------------

func Nil(value any, msg string) {
	if value != nil && !reflect.ValueOf(value).IsNil() {
		panic(msg)
	}
}

func NotNil(value any, msg string) {
	if value == nil || reflect.ValueOf(value).IsNil() {
		panic(msg)
	}
}

// ----------------------------------------------------------------------------
// Order
// ----------------------------------------------------------------------------

func LessThan[T constraints.Ordered](a, b T, msg string) {
	if a >= b {
		panic(msg)
	}
}

func LessThanFunc[T any](a, b T, fn func(a, b T) int, msg string) {
	if fn(a, b) >= 0 {
		panic(msg)
	}
}

func LessThanEq[T constraints.Ordered](a, b T, msg string) {
	if a > b {
		panic(msg)
	}
}

func LessThanEqFunc[T any](a, b T, fn func(a, b T) int, msg string) {
	if fn(a, b) > 0 {
		panic(msg)
	}
}

func GreaterThan[T constraints.Ordered](a, b T, msg string) {
	if a <= b {
		panic(msg)
	}
}

func GreaterThanFunc[T any](a, b T, fn func(a, b T) int, msg string) {
	if fn(a, b) <= 0 {
		panic(msg)
	}
}

func GreaterThanEq[T constraints.Ordered](a, b T, msg string) {
	if a < b {
		panic(msg)
	}
}

func GreaterThanEqFunc[T any](a, b T, fn func(a, b T) int, msg string) {
	if fn(a, b) < 0 {
		panic(msg)
	}
}

// ----------------------------------------------------------------------------
// Zero
// ----------------------------------------------------------------------------

func Zero[T comparable](value T, msg string) {
	var zero T
	if value != zero {
		panic(msg)
	}
}

func NotZero[T comparable](value T, msg string) {
	var zero T
	if value == zero {
		panic(msg)
	}
}
