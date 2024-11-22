package ext

func Ptr[T any](v T) *T {
	return &v
}

func NilPtr[T any]() *T {
	return nil
}
