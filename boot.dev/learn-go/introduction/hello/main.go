package main

import "fmt"

// Hello returns a "hello world" string.
func Hello() string {
	return "hello world"
}

func main() {
	// single-line comments start with "//"
	// comments are just for documentation - they don't execute
	fmt.Println(Hello())
}
