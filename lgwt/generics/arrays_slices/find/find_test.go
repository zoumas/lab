package find_test

import (
	"strings"
	"testing"

	"github.com/zoumas/lab/lgwt/generics/arrays_slices/find"
	"github.com/zoumas/lab/lgwt/generics/assert"
)

func TestFind(t *testing.T) {
	t.Run("the first even number", func(t *testing.T) {
		xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := find.Find(xs, func(x int) bool {
			return x%2 == 0
		})

		assert.True(t, found)
		assert.Equal(t, firstEvenNumber, 2)
	})

	t.Run("the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Benk"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		king, found := find.Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		assert.True(t, found)
		assert.Equal(t, king, Person{Name: "Chris James"})
	})
}

type Person struct {
	Name string
}
