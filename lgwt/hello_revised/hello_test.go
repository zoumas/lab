package hello_test

import (
	"fmt"
	"testing"

	hello "github.com/zoumas/lab/lgwt/hello_revised"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to specified recipient", func(t *testing.T) {
		const want = "Hello Ilias"

		got := hello.Hello("Ilias", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string argument defaults to 'World'", func(t *testing.T) {
		const want = "Hello World"

		got := hello.Hello("", "")

		assertCorrectMessage(t, got, want)
	})

	t.Run("language support for", func(t *testing.T) {
		tests := []struct {
			language string
			name     string
			want     string
		}{
			{hello.Spanish, "Elodie", "Hola Elodie"},
			{hello.French, "Gabriel", "Bonjour Gabriel"},
			{hello.Japanese, "矢口", "こんにちは 矢口"},
		}

		for _, tt := range tests {
			t.Run(tt.language, func(t *testing.T) {
				assertCorrectMessage(t, hello.Hello(tt.name, tt.language), tt.want)
			})
		}
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}

func ExampleHello() {
	fmt.Println(hello.Hello("Ilias", ""))
	// Output: Hello Ilias
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello.Hello("Gabriel", hello.Spanish)
	}
}
