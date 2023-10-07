package main

import "fmt"

func main() {
	// A function is a block of code that does one type of things
	// main() function is also a function
	/*
	 * func [function name]() {
	 *    [instructions]
	 * }
	 */
	greeting("Tommy")
	g := greeting
	fmt.Printf("Function type: %T\n", g)
	g("Tommy")
	greetingMultiple("Tommy", "Bobby", "Andy")
	functionAsParam(greeting, "Tommy")
	fmt.Println(add(1, 2))
	result, ok := add(-1, 2)
	fmt.Println(result, ok)

	fmt.Println(addAsReturnType()(1, 2))

	// Anonymous function
	af := func() { fmt.Println("Hello Anonymous") }
	fmt.Println("Hello Non-Anonymous")
	af()
	func() { fmt.Println("Hello True Anonymous") }()

	fmt.Println(recursion(10))
}

// Recursion (fibonacci)
func recursion(i int) int {
	if i < 2 {
		return i
	}
	return recursion(i-1) + recursion(i-2)
}

func addAsReturnType() func(int, int) (int, bool) {
	return add
}

func add(a int, b int) (int, bool) {
	if a < 0 || b < 0 {
		return 0, false
	}
	return a + b, true
}

func functionAsParam(f func(string), subject string) {
	f(subject)
}

func greeting(subject string) {
	fmt.Println("Hello", subject)
}

func greetingMultiple(subjects ...string) {
	for _, subject := range subjects {
		fmt.Println("Hello", subject)
	}
}
