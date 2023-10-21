package sum

func Reduce[T any](f func(x, y T) T, s []T, initial T) T {
	a := initial
	for _, v := range s {
		a = f(a, v)
	}
	return a
}

func add(x, y int) int {
	return x + y
}

// Sum takes a slice of integers and returns the sum of all elements.
func Sum(n []int) int {
	return Reduce(add, n, 0)
} // 1_793_970 ns => 0.00179397 sec

// Sum takes a slice of integers and returns the sum of all elements.
func sum(n []int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
} // 1_170_719 ns => 0.001170719 sec

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
		}

		tail := xs[1:]
		return append(acc, Sum(tail))
	}

	return Reduce(sumTail, ns, []int{})
}

// SumAllTails works like SumAll but operates on the tails of the slices.
// The tail of a slice is all the elements apart from the first (the head).
func sumAllTails(ns ...[]int) []int {
	var sums []int
	for _, v := range ns {
		if len(v) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := v[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
