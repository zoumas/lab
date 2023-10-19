package shapes_test

import (
	"fmt"
	"testing"

	"github.com/zoumas/lab/lgwt/structs_methods_interfaces/shapes"
)

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := shapes.Rectangle{
			Width:  10.0,
			Height: 10.0,
		}
		want := 40.0

		got := rectangle.Perimeter()

		if got != want {
			t.Errorf("\ngot:\n%.2f\nwant:\n%.2f", got, want)
		}
	})
}

func ExampleRectangle_Perimeter() {
	rectangle := shapes.Rectangle{
		Width:  10.0,
		Height: 10.0,
	}

	perimeter := rectangle.Perimeter()

	fmt.Printf("%.1f\n", perimeter)
	// Output: 40.0
}

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		shape shapes.Shape
		area  float64
	}{
		{name: "Rectangle", shape: shapes.Rectangle{Width: 12, Height: 6}, area: 72},
		{name: "Circle", shape: shapes.Circle{Radius: 10}, area: 314.1592653589793},
		{name: "Triangle", shape: shapes.Triangle{Base: 12, Height: 6}, area: 36},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertArea(t, tt.shape, tt.area)
		})
	}
}

func assertArea(t testing.TB, shape shapes.Shape, want float64) {
	t.Helper()

	if got := shape.Area(); got != want {
		t.Errorf("\ngot:\n%g\nwant:\n%g\ngiven:\n%#v", got, want, shape)
	}
}

func ExampleRectangle_Area() {
	rect := shapes.Rectangle{
		Width:  12.0,
		Height: 6.0,
	}

	area := rect.Area()

	fmt.Printf("%.1f\n", area)
	// Output: 72.0
}

func ExampleCircle_Area() {
	circle := shapes.Circle{
		Radius: 10.0,
	}

	area := circle.Area()

	fmt.Printf("%g\n", area)
	// Output: 314.1592653589793
}

func ExampleTriangle_Area() {
	triangle := shapes.Triangle{
		Base:   12.0,
		Height: 6.0,
	}

	area := triangle.Area()

	fmt.Printf("%.1f\n", area)
	// Output: 36.0
}
