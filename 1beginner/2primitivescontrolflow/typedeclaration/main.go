package main

import "fmt"

func main() {
	// Type declaration is simply renaming an existing type
	type integer int // type [name of new type] [old type]
	var i integer
	fmt.Println(i)

	// Why we need it
	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius
	c = celsius((f - 32) * 5 / 9)
	fmt.Println(f, c)
}
