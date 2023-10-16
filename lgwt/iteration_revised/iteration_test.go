package iteration_test

import (
	"fmt"
	"strings"
	"testing"

	iteration "github.com/zoumas/lab/lgwt/iteration_revised"
)

func TestRepeat(t *testing.T) {
	assertRepeated(t, iteration.Repeat("a", 5), "aaaaa")

	t.Run("behaves like strings.Repeat", func(t *testing.T) {
		assertRepeated(t, iteration.Repeat("ho", 3), strings.Repeat("ho", 3))
	})
}

func assertRepeated(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}

func ExampleRepeat() {
	repeated := "ba" + iteration.Repeat("na", 2)
	fmt.Println(repeated)

	// Output:
	// banana
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
