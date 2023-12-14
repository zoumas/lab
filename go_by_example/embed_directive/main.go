package main

import (
	_ "embed"
	"fmt"
)

//go:embed input
var file string

func main() {
	fmt.Println(file)
}
