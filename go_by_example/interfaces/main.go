package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func measure(shape Shape) {
	fmt.Println(shape)
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}

	measure(r)
	measure(c)
}
