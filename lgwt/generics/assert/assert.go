package assert

import "testing"

func Equal[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%+v\nwant:\n%+v", got, want)
	}
}

func NotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got == want {
		t.Errorf("\ngot:\n%+v\nwant:\n%+v", got, want)
	}
}

func True(t *testing.T, got bool) {
	t.Helper()

	if !got {
		t.Errorf("\ngot: %v\nwant: %v", got, true)
	}
}

func False(t *testing.T, got bool) {
	t.Helper()

	if got {
		t.Errorf("\ngot: %v\nwant: %v", got, false)
	}
}
