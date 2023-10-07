package main

import "fmt"

func main() {
	// Variable is a placeholder of some value subject to change
	// Camel case as naming scheme, ie thisIsAVariable

	// Declaration
	var integerOne int // var [name of variable] [type of the variable] (defaults to 0)
	fmt.Println(integerOne)

	// Assignment & reassignment
	integerOne = 1 // [name of the declared variable] = [value]
	fmt.Println(integerOne)
	integerOne = 2

	// Declaration + assignment
	var i2 = 1
	i3 := 2 // Preferred
	fmt.Println(i2, i3)

	// Multiple variables
	var i4, i5, i6 int
	i4, i5, i6 = 4, 5, 6
	fmt.Println(i4, i5, i6)
	var i7, i8 = 7, 8
	i9, i10 := 9, 10
	fmt.Println(i7, i8, i9, i10)

	var (
		i11 = 11
		i12 = 12
	)
	fmt.Println(i11, i12)
}
