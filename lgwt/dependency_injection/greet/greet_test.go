package greet_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/zoumas/lab/lgwt/dependency_injection/greet"
)

func ExampleGreet() {
	greet.Greet(os.Stdout, "World")
	// Output: Hello, World
}

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}

	greet.Greet(buffer, "Ilias")

	got := buffer.String()
	const want = "Hello, Ilias"

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}
