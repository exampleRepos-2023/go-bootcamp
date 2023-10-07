package main

import "fmt"

func main() {
	// Panic is for stopping the ordinary flow of control in disastrous conditions
	// Recover is for handling such panic's
	defer cleanup()

	const one = 0
	if one != 1 {
		panic("This is impossible!")
	}

}

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Panic is recovered!", r)
	}
}
