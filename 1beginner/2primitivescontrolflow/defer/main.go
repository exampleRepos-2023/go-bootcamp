package main

import "fmt"

func main() {
	// Defer defers a function's execution to the end of the current function
	defer close()
	open()

	fmt.Println(recursion(6))

}

func open() {
	fmt.Println("Opening file")
}

func close() {
	fmt.Println("Closing file")
}

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func third() {
	fmt.Println("Third")
}

// Recursion (fibonacci)
func recursion(i int) int {
	if i < 2 {
		return i
	}
	return recursion(i-1) + recursion(i-2)
}
