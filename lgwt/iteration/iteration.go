package iteration

import "strings"

// Repeat returns a new string consisting of count copies of the string s.
func Repeat(s string, count int) string {
	return RepeatBuilder(s, count)
}

func RepeatConcat(s string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += s
	}
	return repeated
}

func RepeatBuilder(s string, count int) string {
	var b strings.Builder
	for i := 0; i < count; i++ {
		b.WriteString(s)
	}
	return b.String()
}
