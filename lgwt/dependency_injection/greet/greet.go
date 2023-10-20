package greet

import (
	"fmt"
	"io"
)

// Greet writes a greeting to a specifient recipient through an io.Writer
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
