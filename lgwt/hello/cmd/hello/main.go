package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zoumas/lab/lgwt/hello"
)

func main() {
	var name string

	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	} else {
		name = os.Getenv("USER")
	}

	fmt.Println(hello.Hello(name, ""))
}
