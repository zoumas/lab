package main

import "fmt"

func Sum(ns ...int) int {
	var sum int
	for _, v := range ns {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(Sum())
	fmt.Println(Sum(1))
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1, 2, 3))

	ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Sum(ns...))
}
