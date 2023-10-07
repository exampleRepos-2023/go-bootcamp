package main

import (
	"fmt"
	"gitlab.com/golangdojo/bootcamp/2intermediate/3testing/mocking/ninja"
	"gitlab.com/golangdojo/bootcamp/2intermediate/3testing/mocking/secretgen"
)

func main() {
	sg := secretgen.SecretGen{}
	tommy := ninja.Ninja{
		Name:      "Tommy",
		SecretGen: sg,
	}
	secretGreeting := tommy.Greeting()
	fmt.Println(secretGreeting)
}
