package hello_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/hello"
)

func TestHello(t *testing.T) {
	got := hello.Hello("Ilias")
	want := "Hello, Ilias"

	if got != want {
		t.Errorf("\ngot %q\nwant %q", got, want)
	}
}
