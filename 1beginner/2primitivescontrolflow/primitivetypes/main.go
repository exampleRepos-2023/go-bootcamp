package main

import "fmt"

func main() {
	/*
	 * Categories of primitive types
	 * 1. Numeric - int, float64, etc.
	 * 2. Boolean - bool
	 * 3. Alphabetic - byte, rune, string
	 */

	// Numeric
	var integer int             // int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr,
	var floatingPoint float64   // float32
	var complexNumber complex64 // complex128
	fmt.Println(integer, floatingPoint, complexNumber)

	// Boolean
	var boolean bool
	fmt.Println(boolean)

	// Alphabetic
	var character byte     // 'A'
	var utf8Character rune // '中'
	var sentence string    // "Hello World!"
	fmt.Println(character, utf8Character, sentence)
	character = 'A'
	utf8Character = '中'
	sentence = "Hello World!"
	fmt.Println(character, utf8Character, sentence)
	fmt.Println(string(character), string(utf8Character), sentence)
	fmt.Println(sentence[0], sentence[1])
	fmt.Println(string(sentence[0]), string(sentence[1]))
	//sentence[0] = 'c'
	sentence = "Hello Golang Dojo!\n"
	sentence = `Hello Golang Dojo!\n`
	fmt.Println(sentence)
}
