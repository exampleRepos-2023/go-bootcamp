package main

import (
	"bytes"
	"fmt"
)

func main() {
	// make() & new()
	// built-in functions to allocate memory for a given target type T
	// make(T, args) & new(T)

	// Data Types
	// make() only works with slices, maps, and channels & new() works with any
	s := make([]int, 0)
	m := make(map[int]int)
	c := make(chan bool)
	//_ := make(int)
	fmt.Println(s, m, c)
	i := new(int)
	fmt.Println(i)

	// Returned Types
	// make() returns the target type & new() returns the address
	var v map[int]int = make(map[int]int)
	var p *map[int]int = new(map[int]int)
	fmt.Println(v, p)

	// Memory Initialization
	// make() initializes memory & new() zeros memory
	var mMake map[int]int = make(map[int]int) // Initializes a map
	//var mNew *map[int]int = new(map[int]int)  // Zeros a map pointer to nil
	mMake[0] = 1 // Runs
	//(*mNew)[0] = 1                            // Crashes

	// Zero Value
	// Sometimes new()'s zero values are useful rather than just nil
	// The zero value for Buffer is an empty buffer ready to use.
	// type Buffer struct {
	//     buf      []byte
	//	   off      int
	//	   lastRead readOp
	// }
	b1 := new(bytes.Buffer) // Usage won't cause a crash
	var b2 bytes.Buffer     // Same as above besides the variable type
	fmt.Println(b1.Len())
	fmt.Println(b2.Len())

	// make() - slices, maps, and channels
	// new() - easier to create primitive pointers + no full type initialization needed
}
