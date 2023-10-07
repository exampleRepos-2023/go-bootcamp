package main

import "fmt"

func main() {
	// If statements control conditional branching

	if i := 0; i > 0 {
		fmt.Println("Greater than 0")
	} else if i < 0 {
		fmt.Println("Less than 0")
	} else {
		fmt.Println("Equals to 0")
	}
	if j := 1; j > 0 {
		fmt.Println("Greater than 0")
	}

	fmt.Println("Hello World")
}
