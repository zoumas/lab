package sum_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/zoumas/lab/lgwt/generics/reduce/sum"
)

func TestReduce(t *testing.T) {
	t.Run("product of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		assertEqual(t, sum.Reduce(multiply, []int{1, 2, 3, 4, 5}, 1), 120)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concat := func(x, y string) string {
			return x + y
		}

		assertEqual(t, sum.Reduce(concat, []string{"a", "b", "c"}, ""), "abc")
	})
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%#v\nwant:\n%#v", got, want)
	}
}

func TestSum(t *testing.T) {
	assertSum := func(t testing.TB, got, want int, given []int) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot:\n%d\nwant:\n%d\ngiven:\n%v", got, want, given)
		}
	}

	t.Run("array of 5 integers", func(t *testing.T) {
		n := [5]int{1, 2, 3, 4, 5}

		const want = 15
		got := sum.Sum(n[:]) // get a slice from an array

		assertSum(t, got, want, n[:])
	})

	t.Run("slices", func(t *testing.T) {
		n := []int{1, 2, 3}

		want := 6
		got := sum.Sum(n)

		assertSum(t, got, want, n)
	})

	t.Run("nil slice", func(t *testing.T) {
		var n []int

		want := 0
		got := sum.Sum(n)

		assertSum(t, got, want, n)
	})
}

func BenchmarkSum(b *testing.B) {
	s := make([]int, 1e9)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum.Sum(s)
	}
}

func ExampleSum() {
	given := []int{2, 3, 5}
	got := sum.Sum(given)
	fmt.Println(got)
	// Output: 10
}

func TestSumAll(t *testing.T) {
	given := [][]int{{1, 2}, {0, 9}, nil}
	want := []int{3, 9, 0}
	got := sum.SumAll(given...)

	assertSums(t, got, want, given)
}

func BenchmarkSumAllTails(b *testing.B) {
	xs := [][]int{{1, 2, 3, 4}, {0, 9}, nil}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum.SumAllTails(xs...)
	}
}

func ExampleSumAll() {
	given := [][]int{{1, 2}, {3, 4}, {5, 6}}
	got := sum.SumAll(given...)
	fmt.Println(got)
	// Output: [3 7 11]
}

func TestSumAllTails(t *testing.T) {
	given := [][]int{{1, 2}, {0, 9}, {1, 2, 3}, {}, nil}
	want := []int{2, 9, 5, 0, 0}
	got := sum.SumAllTails(given...)

	assertSums(t, got, want, given)
}

func ExampleSumAllTails() {
	given := [][]int{{1, 2, 3}, {0, 9, 1, 2}, {2, 4, 6}, {}, nil}
	got := sum.SumAllTails(given...)
	fmt.Println(got)
	// Output: [5 12 10 0 0]
}

func assertSums(t testing.TB, got, want []int, given [][]int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("\ngot:\n%+v\nwant:\n%+v\ngiven:\n%+v", got, want, given)
	}
}
