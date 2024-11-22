package guard

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

func Equal[T any](a, b T) {
	if !reflect.DeepEqual(a, b) {
		panic("values are not equal")
	}
}

func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Ptr:
		if v.IsNil() {
			return true
		}
		deref := v.Elem().Interface()
		return isEmpty(deref)
	default:
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(value, zero.Interface())
	}
}

func NotEmpty(value any) {
	if isEmpty(value) {
		panic("value is empty")
	}
}

func NilOrNotEmpty(value any) {
	if value != nil && isEmpty(value) {
		panic("value is empty")
	}
}

func Each[T any](s []T, fn func(v any)) {
	for _, v := range s {
		fn(v)
	}
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case
		reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		if v.IsNil() {
			return true
		}
	}

	return false
}

func NotNil(value any) {
	if isNil(value) {
		panic("value is nil")
	}
}

func isZero(value any) bool {
	if value == nil {
		return true
	}

	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func NotZero(value any) {
	if isZero(value) {
		panic("value is zero")
	}
}

func GT[T constraints.Ordered](value T, target T) {
	if value <= target {
		panic(fmt.Sprintf("%v is not greater than %v:", value, target))
	}
}

func GTE[T constraints.Ordered](value T, target T) {
	if value < target {
		panic(fmt.Sprintf("%v is not greater than or equal to %v", value, target))
	}
}

func LT[T constraints.Ordered](value T, target T) {
	if value >= target {
		panic(fmt.Sprintf("%v is not less than %v", value, target))
	}
}

func LTE[T constraints.Ordered](value T, target T) {
	if value > target {
		panic(fmt.Sprintf("%v is not less than or equal to %v", value, target))
	}
}

func True(value bool) {
	if !value {
		panic("value is false when expected true")
	}
}

func False(value bool) {
	if value {
		panic("value is true when expected false")
	}
}

func Xor(a, b bool) {
	if !((a || b) && !(a && b)) {
		panic("values are not exclusive")
	}
}
