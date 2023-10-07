package main

import "fmt"

func main() {
	// nil is the absence of value
	// If a variable is nil, there's an absence of value for that variable

	// Why do we need nil? Why do we need to express the absence of value?
	var i int
	var f float64
	var s string
	fmt.Printf("i=[%d], f=[%f], s=[%s]\n", i, f, s)

	var p *int // What would the default be?
	if i == 0 {
		fmt.Println("i is defaulted to 0")
	}
	if p == nil {
		fmt.Println(p)
	}
	
	/* nil is default zero value for
	 * 		-pointers
	 *		-interfaces
	 *		-slices
	 *		-channels
	 *		-maps
	 *		-functions
	 */
}
