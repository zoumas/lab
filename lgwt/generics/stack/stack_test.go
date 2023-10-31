package stack_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/generics/assert"
	"github.com/zoumas/lab/lgwt/generics/stack"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		s := new(stack.Stack[int])

		assert.True(t, s.IsEmpty())

		s.Push(123)
		assert.False(t, s.IsEmpty())

		s.Push(456)
		v, _ := s.Pop()
		assert.Equal(t, v, 456)
		v, _ = s.Pop()
		assert.Equal(t, v, 123)
		assert.True(t, s.IsEmpty())

		s.Push(1)
		s.Push(2)

		first, _ := s.Pop()
		second, _ := s.Pop()
		assert.Equal(t, first+second, 3)
	})

	// duplication
	t.Run("string stack", func(t *testing.T) {
		s := new(stack.Stack[string])

		assert.True(t, s.IsEmpty())

		s.Push("Bob")
		assert.False(t, s.IsEmpty())

		s.Push("Alice")
		v, _ := s.Pop()
		assert.Equal(t, v, "Alice")
		v, _ = s.Pop()
		assert.Equal(t, v, "Bob")
		assert.True(t, s.IsEmpty())
	})
}
