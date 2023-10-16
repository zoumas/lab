package iteration

import "strings"

// Repeat repeats s count times.
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
	var builder strings.Builder
	for i := 0; i < count; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}
