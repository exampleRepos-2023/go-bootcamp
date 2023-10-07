package main

import "fmt"

type ninja struct { // type [name of struct] struct {}
	name  string
	level int
}

// Methods of a struct are customized behaviors to that struct
func (ninja) greet() {
	fmt.Println("Hello")
}

func (n *ninja) greetWithName() {
	n.level++
	fmt.Println("Hello, I'm", n.name)
}

type dojo struct {
	name   string
	ninjas []ninja
}

func main() {
	// A struct is a user-defined type that represent a collection of fields

	// A ninja has a name and a level
	var name string
	var level int
	fmt.Println(name, level)

	// Instantiation
	tommy := ninja{}
	fmt.Println(tommy)
	tommy = ninja{"Tommy", 9001}
	fmt.Println(tommy)
	tommy = ninja{name: "Tommy", level: 9001}
	fmt.Println(tommy)

	// Individual fields
	fmt.Println(tommy.name, tommy.level)
	tommy.name = "Tommy Jr"
	tommy.level++
	fmt.Println(tommy.name, tommy.level)

	ninjas := []ninja{tommy}
	golangDojo := dojo{"Golang Dojo", ninjas}
	fmt.Println(golangDojo)

	// Individual fields
	golangDojo.name = "Golang Dojo Sr"
	golangDojo.ninjas[0].name = "Tommy Sr"
	fmt.Println(golangDojo)

	// Reference type (slice, pointer...) vs non-reference type (string, int...)
	fmt.Println(tommy)
	fmt.Println(ninjas)

	// Methods
	tommy.greet()
	tommy.greetWithName()
	fmt.Println(tommy.level)
}
