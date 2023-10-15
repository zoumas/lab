package adder_test

import (
	"fmt"
	"testing"

	"github.com/zoumas/lab/lgwt/integers/adder"
)

func TestAdd(t *testing.T) {
	const want = 4

	got := adder.Add(2, 2)

	if got != want {
		t.Errorf("\ngot:\n%d\nwant:\n%d", got, want)
	}
}

func ExampleAdd() {
	sum := adder.Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
