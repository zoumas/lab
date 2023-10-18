package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)

	v3 := m["k3"]
	fmt.Println(v3)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	_, ok := m["k2"]
	fmt.Println(ok)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	if maps.Equal(n, map[string]int{"foo": 1, "bar": 2}) {
		fmt.Println("maps are equal")
	}
}
