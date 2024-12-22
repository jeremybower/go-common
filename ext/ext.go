package ext

func If[T any](condition bool, value T, elseValue T) T {
	if condition {
		return value
	}
	return elseValue
}

func Ptr[T any](v T) *T {
	return &v
}

func NilPtr[T any]() *T {
	return nil
}
