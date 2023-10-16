package repeat_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/zoumas/lab/lgwt/iteration/repeat"
)

func TestRepeat(t *testing.T) {
	assertRepeated := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
		}
	}

	t.Run("repeat 5 times", func(t *testing.T) {
		const want = "aaaaa"

		got := repeat.Repeat("a", 5)

		assertRepeated(t, got, want)
	})

	t.Run("behaves like strings.Repeat", func(t *testing.T) {
		want := "ba" + strings.Repeat("na", 2)

		got := "ba" + repeat.Repeat("na", 2)

		assertRepeated(t, got, want)
	})
}

func ExampleRepeat() {
	fmt.Println("ba" + repeat.Repeat("na", 2))
	// Output:
	// banana
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat.Repeat("a", 5)
	}
}
