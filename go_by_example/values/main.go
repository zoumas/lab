package main

import (
	"fmt"
)

func main() {
	// string concatenation (joining)
	fmt.Println("go" + "lang")

	// basic arithmetic
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// logic operations
	fmt.Println(true && false) // logical and
	fmt.Println(true || false) // logical or
	fmt.Println(!true)         // logical not - inverter
}
