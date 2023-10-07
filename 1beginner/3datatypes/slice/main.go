package main

import "fmt"

func main() {
	// A slice is a growing-sized collection of elements of a single type

	var s []int // No size specified
	if s == nil {
		fmt.Println("nil is the default zero value of a slice")
	}
	s = []int{0, 1}
	fmt.Printf("s=%v, type=%T\n", s, s)
	s = []int{0, 1, 2}
	fmt.Printf("s=%v, type=%T\n", s, s)

	s = append(s, 3, 4, 5) // variadic
	fmt.Println(s)

	// Individual elements
	fmt.Println(s[0])
	s[0] = 10
	fmt.Println(s[0])

	// Iterations
	for index, value := range s {
		fmt.Printf("%d: %d | ", index, value)
	}
	fmt.Println()

	// len() - number of elements we added to the slice
	fmt.Println(len(s))

	// cap() - number of elements we *can* add before having to re-allocate more space
	fmt.Println(cap(s))
	s = append(s, 1, 2, 3)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// Pre-allocation
	s = make([]int, 5)
	fmt.Println(s)
	fmt.Println(cap(s))

	// Slicing
	s = []int{0, 1, 2, 3}
	s = s[0:2]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)

	// Multi-dimensional
	s2D := [][]int{{}}
	fmt.Println(s2D)
	s2D = [][]int{{0, 1}, {2, 3}}
	s3D := [][][]int{{{0, 1}, {2, 3}}, {{0, 1}, {2, 3}}}
	fmt.Println(s2D)
	fmt.Println(s3D)

	// Again, same for other types
	sString := []string{"Tommy", "Bobby"}
	fmt.Println(sString)
}
