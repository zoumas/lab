package hello_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/hello"
)

func TestHello(t *testing.T) {
	t.Run("say hello to specified recipient", func(t *testing.T) {
		const want = "Hello Ilias"

		got := hello.Hello("Ilias", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		const want = "Hello World"

		got := hello.Hello("", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in different languages", func(t *testing.T) {
		tests := []struct {
			language string
			name     string
			want     string
		}{
			{"Spanish", "Elodie", "Hola Elodie"},
			{"French", "Gabriel", "Bonjour Gabriel"},
			{"Japanese", "矢口八虎", "こんにちは 矢口八虎"},
		}

		for _, test := range tests {
			t.Run(test.language, func(t *testing.T) {
				assertCorrectMessage(t, hello.Hello(test.name, test.language), test.want)
			})
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot %q\nwant %q", got, want)
	}
}
