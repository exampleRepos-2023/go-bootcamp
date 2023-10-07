package main

import "fmt"

func main() {
	// Arrays are a fix-sized collection of elements of a single type

	i, j, k := 0, 1, 2
	// 10
	fmt.Println(i, j, k)

	intArray := [10]int{} // [(number of element)]int[]
	fmt.Println(intArray)
	intArray = [10]int{0, 1, 2}
	intArray = [10]int{0: 0, 1: 1, 2: 2}
	intArray = [10]int{0: 0, 1: 1, 2}
	fmt.Println(intArray)
	//intArray = [9]int{}
	fmt.Printf("%T\n", intArray)

	for index, value := range intArray {
		fmt.Printf("%d: %d | ", index, value)
	}
	fmt.Println()

	intArrayImplicit := [...]int{0, 1}
	fmt.Println(intArrayImplicit)
	fmt.Printf("%T\n", intArrayImplicit)

	fmt.Println(intArray[0])
	intArray[0] = 10
	fmt.Println(intArray[0])

	intArrayCopy := intArray
	fmt.Println(intArrayCopy)
	intArrayCopy[0] = 0
	fmt.Println(intArray[0])

	intArray2D := [2][2]int{{0, 1}, {2, 3}}
	fmt.Println(intArray2D)
	intArray3D := [2][2][2]int{{{0, 1}, {2, 3}}, {{0, 1}, {2, 3}}}
	fmt.Println(intArray3D)

	stringArray := [...]string{"Tommy", "Bobby"}
	fmt.Println(stringArray)
}
