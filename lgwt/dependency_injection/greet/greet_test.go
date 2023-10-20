package greet_test

import (
	"bytes"
	"testing"

	"github.com/zoumas/lab/lgwt/dependency_injection/greet"
)

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}

	greet.Greet(buffer, "Ilias")

	want := "Hello, Ilias"
	got := buffer.String()

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}
