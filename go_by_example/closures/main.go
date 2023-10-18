package main

import "fmt"

func Counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	c := Counter()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	fmt.Println()
	c2 := Counter()

	fmt.Println(c())
	fmt.Println(c2())
}
