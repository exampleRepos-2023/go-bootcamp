package main

import "fmt"

func main() {
	// Address & Pointer (of a variable)
	// A pointer variable is a placeholder that holds the memory address of a variable
	// We can update the value of the variable using the pointer to / address of that variable
	var integer int = 1
	var integerPointer *int = &integer
	fmt.Printf("integer=%d, type=%T\nintegerPointer=%d, type=%T\n",
		integer, integer, integerPointer, integerPointer)

	// We use "*" 1) as a type, 2) as dereference of the pointer variable
	// dereference of a pointer variable is to get the original variable
	// (integerPointer points to integer)
	// *integerPointer <=> integer
	fmt.Printf("*integerPointer=%d, type=%T\n",
		*integerPointer, *integerPointer)

	// You change the content/value of the variable by changing
	// the content/value of the pointer variable's dereference
	*integerPointer = 2 // same as integer = 2
	fmt.Printf("integer=%d, type=%T\nintegerPointer=%d, type=%T\n",
		integer, integer, integerPointer, integerPointer)

	// Use case
	var i int = 1
	update(&i)
	fmt.Printf("i=%d\n", i)
}

func update(iPointer *int) {
	*iPointer = 2
}
