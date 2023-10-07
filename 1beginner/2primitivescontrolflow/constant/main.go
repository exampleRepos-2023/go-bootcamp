package main

import "fmt"

func main() {
	// Constants are placeholders with fixed values
	const noon = 12
	const firstProgram string = "Hello World!"
	const (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(noon, firstProgram, a, b, c)

	const (
		d = 1
		e
		f
	)
	fmt.Println("bosluklu", d, e, f)

	const (
		zero = iota
		one
		two
		three
		// ...
	)
	fmt.Println(zero, one, two, three)

	const (
		five = iota + 5
		six
		seven
	)
	fmt.Println(five, six, seven)
}
