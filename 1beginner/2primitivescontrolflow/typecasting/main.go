package main

import "fmt"

func main() {
	// Type casting is converting one type to another, if they are convertable
	var i int
	var f float64
	f = float64(i)
	i = int(f)
	fmt.Printf("i=%d %T, f=%f %T\n", i, i, f, f)

	// int takes the floor of float when casting
	f = 1.7
	i = int(f)
	fmt.Printf("i=%d %T, f=%f %T\n", i, i, f, f)

	// No casting between non-convertable types
	//s := "Hello World!"
	//c := byte(s)
}
