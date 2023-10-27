package walk_test

import (
	"slices"
	"testing"

	"github.com/zoumas/lab/lgwt/reflection/walk"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"plain string",
			"Chris",
			[]string{"Chris"},
		},
		{
			"struct with one string fields",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointer to struct",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, cs := range cases {
		t.Run(cs.Name, func(t *testing.T) {
			want := cs.ExpectedCalls

			var got []string
			walk.Walk(cs.Input, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, want) {
				t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
			}
		})
	}

	t.Run("maps", func(t *testing.T) {
		m := map[string]string{
			"Name": "Chris",
			"City": "London",
		}

		var got []string
		walk.Walk(m, func(input string) {
			got = append(got, input)
		})

		for _, v := range m {
			assertMapContains(t, got, v)
		}
	})

	t.Run("channels", func(t *testing.T) {
		c := make(chan Profile)

		go func() {
			c <- Profile{33, "Berlin"}
			c <- Profile{34, "Katowice"}
			close(c)
		}()
		want := []string{"Berlin", "Katowice"}

		var got []string
		walk.Walk(c, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
		}
	})

	t.Run("functions", func(t *testing.T) {
		f := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		want := []string{"Berlin", "Katowice"}
		var got []string

		walk.Walk(f, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("\ngot:\n%v\nwant:\n%v", got, want)
		}
	})
}

func assertMapContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("\nexpected %v\nto contain %q", haystack, needle)
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
