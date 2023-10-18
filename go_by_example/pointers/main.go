package main

import "fmt"

func zeroByValue(val int) {
	fmt.Printf("Address of value in zeroByValue : %p\n", &val)
	val = 0
}

func zeroByPointer(val *int) {
	fmt.Printf("Address of value in zeroByPointer : %p\n", val)
	*val = 0
}

func main() {
	val := 1
	fmt.Printf("Address of value in main : %p\n", &val)

	zeroByValue(val)
	fmt.Println(val)

	zeroByPointer(&val)
	fmt.Println(val)
}
