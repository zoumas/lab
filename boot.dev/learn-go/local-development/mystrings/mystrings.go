package mystrings

// Reverse reverses a string left to right.
func Reverse(s string) string {
	var reversed string
	for _, v := range s {
		reversed = string(v) + reversed
	}
	return reversed
}
