package main

import "fmt"

func main() {
	/*
	 * Categories of operators
	 * 1. Logical || && !
	 * 2. Mathematical + - * / ^ %
	 * 3. Relational == != > < >= <=
	 * 4. Assignment = += %= ++ --...
	 * 5. Bitwise & |...
	 */

	// Logical
	fmt.Println(true, true || false, true && false, !true)

	// Mathematical
	fmt.Println(1+1, 1-1, 1/1, 1*1, 1%1)

	// Relational
	fmt.Println(0 == 0, 0 != 0, 0 < 0, 0 > 0, 0 <= 0, 0 >= 0)

	// Assignment
	i := 0
	i = 1
	i += 1 // i = i + 1
	i -= 1 // i = i - 1
	i /= 1 // i = i / 1
	i *= 1 // i = i * 1
	i %= 1 // i = i % 1
	i++
	i--

	// Bitwise
	i = i << 1
	i = i >> 1
	i <<= 1 // i = i << 1
	i >>= 1 // i = i >> 1
	fmt.Println(1|0, 1&0)
}
