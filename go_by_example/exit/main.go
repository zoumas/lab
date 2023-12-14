package main

import (
	"fmt"
	"os"
)

// Use os.Exit to immediately exit with a given status code.
// defers will not be run when using os.Exit.

func main() {
	defer fmt.Println("!")
	os.Exit(1)
}
