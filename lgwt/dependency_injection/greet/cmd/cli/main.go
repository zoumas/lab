package main

import (
	"os"

	"github.com/zoumas/lab/lgwt/dependency_injection/greet"
)

func main() {
	name := "World"
	greet.Greet(os.Stdout, name)
}
