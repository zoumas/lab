package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

func main() {
	r := Rectangle{Width: 10, Height: 5}

	fmt.Println("Area:", r.Area())
	fmt.Println("Perimeter:", r.Perimeter())

	rp := &r
	fmt.Println("Area:", rp.Area())
	fmt.Println("Perimeter:", rp.Perimeter())
}
