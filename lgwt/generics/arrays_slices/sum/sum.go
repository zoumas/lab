package sum

func Reduce[A any](s []A, acc func(A, A) A, identity A) A {
	rv := identity
	for _, v := range s {
		rv = acc(rv, v)
	}
	return rv
}

func add(x, y int) int {
	return x + y
}

// Sum takes a slice of integers and returns the sum of all elements.
func Sum(n []int) int {
	return Reduce(n, add, 0)
}

// SumAll takes a variable size of slices of integers
// and returns a new slice where each element is the sum of all elements of the correspoding slice.
func SumAll(ns ...[]int) []int {
	sums := make([]int, len(ns))
	for i, v := range ns {
		sums[i] = Sum(v)
	}
	return sums
}

// SumAllTails works like SumAll but operates on the tails of the slices.
// The tail of a slice is all the elements apart from the first (the head).
func SumAllTails(ns ...[]int) []int {
	sumTail := func(acc, xs []int) []int {
		if len(xs) == 0 {
			return append(acc, 0)
		} else {
			tail := xs[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(ns, sumTail, []int{})
}
