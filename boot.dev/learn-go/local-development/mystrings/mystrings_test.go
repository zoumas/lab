package mystrings_test

import (
	"testing"

	"github.com/zoumas/lab/boot.dev/learn-go/local-development/mystrings"
)

func TestReverse(t *testing.T) {
	const want = "sailI"

	got := mystrings.Reverse("Ilias")

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}
