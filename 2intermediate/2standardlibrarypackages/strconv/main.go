package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Package strconv implements conversions
	// between strings and other basic data types

	i, err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, err, s)

	b, err := strconv.ParseBool("true")
	i64, err := strconv.ParseInt("-42", 10, 64)
	fmt.Println(b, i64)

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(-42, 16))

	var bytes []byte
	bytes = strconv.AppendBool(bytes, true)
	fmt.Println(string(bytes))
	bytes = strconv.AppendInt(bytes, -42, 16)
	fmt.Println(string(bytes))

	// Show "" and escape the escape literals
	fmt.Println("Hello World")
	fmt.Println("Hello \"World\"")
	fmt.Println(strconv.Quote("Hello \"World\""))
}
