package ptr

// Wrap converts a value into a pointer.
func Wrap[T any](v T) *T {
	return &v
}

// Unwrap derefences a pointer, or returns the zero-value of T if the pointer is
// nil.
func Unwrap[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}
