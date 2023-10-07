package main

import (
	"fmt"
)

func main() {
	ok := validate("")
	if !ok {
		fmt.Println("Something went wrong")
	}
	ok = validate(" ")
	if !ok {
		fmt.Println("Something went wrong")
	}

	err := validateWithError("")
	if err != nil {
		fmt.Println("Something went wrong. The reason is:", err.Error())
	}
}

// Error handling
func validate(s string) bool {
	if len(s) == 0 {
		return false
	} else if s == " " {
		return false
	}
	return true
}

func validateWithError(s string) error {
	if len(s) == 0 {
		return myNewError("length is 0")
	} else if s == " " {
		return myNewError("s is a single space")
	}
	return nil
}

// Custom error
type myError struct {
	message string
}

func (err *myError) Error() string {
	return "special error: " + err.message
}

func myNewError(message string) error {
	return &myError{message}
}
