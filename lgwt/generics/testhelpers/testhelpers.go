package testhelpers

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%#v\nwant:\n%#v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got == want {
		t.Errorf("\ngot:\n%#v\nwant:\n%#v", got, want)
	}
}
