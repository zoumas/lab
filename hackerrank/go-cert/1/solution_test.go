package solution_test

import (
	"slices"
	"testing"

	solution "github.com/zoumas/lab/hackerrank/go-cert/1"
)

func TestSolution(t *testing.T) {
	// strArr := []string{"Colorado", "Utah", "Wisconsin", "Oregon"}
	strArr := []string{"a", "ab", "bc", "abc"}
	got := solution.Solution1(strArr)
	// want := []string{"Oregon", "Wisconsin", "Utah", "Colorado"}
	want := []string{"abc", "a", "ab", "bc"}

	if !slices.Equal(got, want) {
		t.Errorf("\ngot: %v\nwant: %v", got, want)
	}
}
