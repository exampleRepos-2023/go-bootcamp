package main

import (
	"fmt"
	"gitlab.com/golangdojo/bootcamp/1beginner/2primitivescontrolflow/scope/is"
)

// Scope of a variable is the accessibility and the lifetime of that variable
// There are 3 levels of scope -- block, unexported & exported

func main() {
	//is.f()
	is.F()
	//fmt.Println(is.i)
	fmt.Println(is.I)
	{
		i := 0
		fmt.Println(i)
	}
	i := 1
	fmt.Println(i)
}
