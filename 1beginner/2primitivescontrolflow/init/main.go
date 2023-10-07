package main

import (
	"fmt"
	"gitlab.com/golangdojo/bootcamp/1beginner/2primitivescontrolflow/init/i"
)

func main() {
	fmt.Println(i.I)
	fmt.Println("Hello World!")
}

func init() {
	fmt.Println("main init")
}
