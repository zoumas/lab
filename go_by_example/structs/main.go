// Go by Example: Structs

// Go's structs are typed collections of fields.
// They're useful for grouping data together to form records.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Person "constructor"
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{Name: "Alice", Age: 30})

	fmt.Println(Person{Name: "Fred"})
	fmt.Println(&Person{Name: "Ann", Age: 40})

	fmt.Println(NewPerson("John", 42))

	s := Person{Name: "Sean", Age: 50}
	fmt.Println(s.Name)

	sp := &s
	fmt.Println(sp.Age)

	sp.Age = 51
	fmt.Println(sp.Age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
