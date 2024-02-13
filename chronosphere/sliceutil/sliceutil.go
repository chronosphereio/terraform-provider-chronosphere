package sliceutil

// Convert a slice of one object into another using a function
func Map[X any, Y any](xs []X, f func(X) Y) []Y {
	if xs == nil {
		return nil
	}
	ys := make([]Y, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Convert a slice of one object into another using a function that returns an error.
func MapErr[X any, Y any](xs []X, f func(X) (Y, error)) ([]Y, error) {
	if xs == nil {
		return nil, nil
	}
	ys := make([]Y, 0, len(xs))
	for _, x := range xs {
		y, err := f(x)
		if err != nil {
			return nil, err
		}
		ys = append(ys, y)
	}
	return ys, nil
}
