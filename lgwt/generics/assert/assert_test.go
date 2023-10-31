package assert_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/generics/assert"
)

func TestAssert(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		assert.Equal(t, 1, 1)
		assert.NotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		assert.Equal(t, "hello", "hello")
		assert.NotEqual(t, "hello", "Grace")
	})

	// compilation errors
	// assert.Equal(t, 1, "1")
	// assert.Equal(t, 1, []int{1})
	t.Run("assert truth", func(t *testing.T) {
		assert.True(t, 1 == 1)
		assert.False(t, "1"+"1" != "11")
	})
}
