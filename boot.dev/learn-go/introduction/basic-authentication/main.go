package main

import "fmt"

func main() {
	var username string = "wagslane"
	var password string = "20947382822"

	// Two strings can be concatenated with the `+` operator.
	// Because Go is strongly typed, it won't allow you to concatenate
	// a string variable with a numeric variable.
	fmt.Println("Authorization: Basic", username+":"+password)
}
