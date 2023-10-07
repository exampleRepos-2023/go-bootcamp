package main

import (
	"gitlab.com/golangdojo/bootcamp/3advanced/1objectorientedprogramming/encapsulation/ninja"
)

func main() {
	n := ninja.New()
	n.Greet()
	//fmt.Println(n.Name)
	//fmt.Println(n.Level)
	//n.Name += " Sr"
	//n.Level += 1
}
