package repeat

// Repeat takes a string s and returns string of s repeated 5 times
func Repeat(s string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += s
	}
	return repeated
}
