package find_test

import (
	"strings"
	"testing"

	"github.com/zoumas/lab/lgwt/generics/find"
)

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := find.Find(xs, func(x int) bool {
			return x%2 == 0
		})

		if !found {
			t.Errorf("should have found number")
		}

		assertEqual(t, firstEvenNumber, 2)
	})

	t.Run("find the best programmer", func(t *testing.T) {
		type Person struct {
			Name string
		}

		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		best, found := find.Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Ken")
		})

		if !found {
			t.Errorf("should have found Person")
		}
		assertEqual(t, best, Person{Name: "Kent Beck"})
	})
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("\ngot:\n%#v\nwant:\n%#v", got, want)
	}
}
