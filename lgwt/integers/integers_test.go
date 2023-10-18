package integers_test

import (
	"fmt"
	"testing"

	"github.com/zoumas/lab/lgwt/integers"
)

func TestAdd(t *testing.T) {
	assertSum(t, integers.Add(2, 2), 4)
}

func ExampleAdd() {
	sum := integers.Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}

func assertSum(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%d\nwant:\n%d", got, want)
	}
}
