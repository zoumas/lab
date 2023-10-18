package iteration_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/zoumas/lab/lgwt/iteration"
)

func TestRepeat(t *testing.T) {
	t.Run("basic usage", func(t *testing.T) {
		const want = "aaaaa"
		got := iteration.Repeat("a", 5)
		assertRepeated(t, got, want)
	})

	t.Run("behaves like strings.Repeat", func(t *testing.T) {
		const s = "ho"
		const count = 3
		assertRepeated(t, iteration.Repeat(s, count), strings.Repeat(s, count))
	})
}

func ExampleRepeat() {
	fmt.Println(iteration.Repeat("ho", 3))
	// Output: hohoho
}

func BenchmarkRepeatConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.RepeatConcat("a", 5)
	}
}

func BenchmarkRepeatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.RepeatBuilder("a", 5)
	}
}

func assertRepeated(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}
