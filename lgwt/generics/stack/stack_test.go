package stack_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/generics/stack"
)

func TestStack(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		s := new(stack.Stack[int])

		// check stack is empty
		assertIsEmpty(t, s)

		// push a value, then check it's not empty
		s.Push(123)
		assertIsNotEmpty(t, s)

		// add a value, pop it back
		s.Push(456)
		v, _ := s.Pop()
		assertEqual(t, v, 456)
		v, _ = s.Pop()
		assertEqual(t, v, 123)
		_, err := s.Pop()
		assertError(t, err, stack.PopErr)
		assertIsEmpty(t, s)
	})

	t.Run("strings", func(t *testing.T) {
	})
}

func assertIsEmpty[T any](t testing.TB, s *stack.Stack[T]) {
	t.Helper()

	if !s.IsEmpty() {
		t.Errorf("stack should be empty")
	}
}

func assertIsNotEmpty[T any](t testing.TB, s *stack.Stack[T]) {
	t.Helper()

	if s.IsEmpty() {
		t.Errorf("stack should not be empty")
	}
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%#v\nwant:\n%#v", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("expected an error")
	}

	if got != want {
		t.Errorf("\ngot:\n%q\nwant:\n%q", got, want)
	}
}
