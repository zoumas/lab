package testhelpers_test

import (
	"testing"

	"github.com/zoumas/lab/lgwt/generics/testhelpers"
)

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		testhelpers.AssertEqual(t, 0, 0)
		testhelpers.AssertNotEqual(t, 0, 1)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		testhelpers.AssertEqual(t, "hello", "hello")
		testhelpers.AssertNotEqual(t, "hello", "Hello")
	})

	// testhelpers.AssertEqual(t, 1, "1")
}
