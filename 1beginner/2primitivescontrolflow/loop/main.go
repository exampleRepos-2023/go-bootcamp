package main

import "fmt"

func main() {
	// Repeat for as long as the condition is true
	for boolean := true; boolean; boolean = false {
		fmt.Println("Hello World!")
	}

	// Range returns key-value pair of each iteration of the loop
	for _, value := range "Hello!" {
		fmt.Println(string(value))
	}

	// Until the desired condition is met
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
