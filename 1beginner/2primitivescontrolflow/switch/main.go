package main

import "fmt"

func main() {
	// Switch statements allows a value or expression to be tested

	// Condition switch
	i := -1
	switch {
	case i < 0:
		fmt.Println("Less than zero")
	case i > 0:
		fmt.Println("Greater than zero")
	default:
		fmt.Println("Equals to zero")
	}

	// Variable-value switch
	i = 3
	switch i {
	case 0:
		fmt.Println("It's zero")
	case 1, 2:
		fmt.Println("It's one or two")
	default:
		fmt.Println("Not sure")
	}

	// Type switch
	//i = 1
	//switch i.(type) {
	//
	//}

}
