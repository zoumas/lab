package main

import (
	"os"

	"github.com/zoumas/lab/lgwt/dependency_injection/greet"
)

func main() {
	name := os.Getenv("USER")
	if name == "" {
		name = "World"
	}
	greet.Greet(os.Stdout, name)
}
