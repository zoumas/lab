package hello_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/hello"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		const want = "Hello, Ilias"

		got := hello.Hello("Ilias")

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		const want = "Hello, World"

		got := hello.Hello("")

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot %q\nwant %q", got, want)
	}
}
