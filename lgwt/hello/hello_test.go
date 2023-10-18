package hello_test

import (
	"fmt"
	"testing"

	"github.com/zoumas/lab/lgwt/hello"
)

func TestHello(t *testing.T) {
	t.Run("greeting a specified recipinet", func(t *testing.T) {
		const want = "Hello Ilias"
		got := hello.Hello("Ilias", "")
		assertGreeting(t, got, want)
	})

	t.Run("returns 'Hello World' when an empty string is supplied", func(t *testing.T) {
		const want = "Hello World"
		got := hello.Hello("", "")
		assertGreeting(t, got, want)
	})

	t.Run("language support", func(t *testing.T) {
		tests := []struct {
			language string
			name     string
			want     string
		}{
			{hello.Spanish, "Elodie", "Hola Elodie"},
			{hello.French, "Gabriel", "Bonjour Gabriel"},
			{hello.Japanese, "矢口八虎", "こんにちは 矢口八虎"},
		}

		for _, tt := range tests {
			t.Run(tt.language, func(t *testing.T) {
				assertGreeting(t, hello.Hello(tt.name, tt.language), tt.want)
			})
		}
	})
}

func ExampleHello() {
	fmt.Println(hello.Hello("Ilias", ""))
	fmt.Println(hello.Hello("", ""))
	// Output:
	// Hello Ilias
	// Hello World
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello.Hello("矢口八虎", hello.Japanese)
	}
}

func assertGreeting(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}
