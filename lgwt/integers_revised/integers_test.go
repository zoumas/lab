package integers_test

import (
	"fmt"
	"testing"

	integers "github.com/zoumas/lab/lgwt/integers_revised"
)

func TestAdd(t *testing.T) {
	assertSum(t, integers.Add(2, 2), 4)
}

func assertSum(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%d\nwant:\n%d", got, want)
	}
}

func ExampleAdd() {
	sum := integers.Add(2, 2)
	fmt.Println(sum)
	// Output: 4
}
