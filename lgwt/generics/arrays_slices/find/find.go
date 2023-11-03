package find

func Find[T any](xs []T, predicate func(T) bool) (value T, found bool) {
	for _, x := range xs {
		if predicate(x) {
			return x, true
		}
	}

	return
}
