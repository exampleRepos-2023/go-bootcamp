package main

import (
	"fmt"
	"strings"
)

func main() {
	// Package strings
	// implements simple functions to manipulate UTF-8 encoded strings
	fmt.Println(strings.Contains("Hello World!", "Hello"))
	fmt.Println(strings.Compare("1", "2"))
	fmt.Println(strings.Compare("2", "2"))
	fmt.Println(strings.Compare("2", "1"))
	fmt.Println(strings.Count("Hello World", "o"))
	fmt.Println(strings.Index("Hello World!", "World"))
	// ...

	// Builder helps build strings incrementally
	sb := strings.Builder{}
	sb.WriteString("Hello ")
	fmt.Println(sb.String())
	sb.WriteString("World!")
	fmt.Println(sb.String())
	sb.WriteRune(rune('中'))
	sb.WriteByte(byte('A'))
	sb.Write([]byte{'B', 'C'})
	fmt.Println(sb.String())
	sb.Reset()
	fmt.Println(sb.String())
	fmt.Println(sb.Len())
	fmt.Println(sb.Cap())
	sb.Grow(10) // 2 * cap + n
	fmt.Println(sb.Len())
	fmt.Println(sb.Cap())
}
