package sum

// Sum takes a slice of integers and returns the sum of all elements.
func Sum(n []int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
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
