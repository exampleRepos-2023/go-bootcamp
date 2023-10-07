package main

import (
	"fmt"
	"os"
)

func main() {
	// Package fmt implements functions for formatted I/O

	// Print statements
	var i int = 1
	var f float64 = 1.5
	var s string = "Hello World!"
	numArgs, err := fmt.Println(i, f, s)
	fmt.Println(numArgs, err)
	fmt.Print("Hello World!\n")
	fmt.Printf("%s %d %f %t %v %v %v %v\n",
		"Hello World!", 1, 0.1, true,
		"Hello World!", 1, 0.1, true)
	/* Common escape literals
	 * %d -> int
	 * %f -> float64
	 * %s -> string
	 * %t -> bool
	 * %T -> type
	 * %v -> auto
	 */
	s = fmt.Sprintf("%d", i)
	fmt.Println(s)
	fmt.Fprintf(os.Stdout, "%v\n", "Hello World!")

	// Scanning
	//numArgs, err = fmt.Scan(&i, &f)
	//fmt.Println("Scan:", i, f, numArgs, err)
	//_, err = fmt.Scanf("%d, %s\n", &i, &s)
	//fmt.Println("Scanf:", i, s, err)
	//_, err = fmt.Fscan(os.Stdin, &i)
	//fmt.Println("Fscan:", i, err)
	_, err = fmt.Fscanf(os.Stdin, "%d, %s\n", &i, &s)
	fmt.Println("Fscanf:", i, s, err)
}
