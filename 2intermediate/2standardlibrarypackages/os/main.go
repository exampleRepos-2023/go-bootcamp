package main

import (
	"fmt"
	"os"
)

func main() {
	// Package os provides an interface to the operating system
	// io, file management, environment variables, etc.

	n, err := os.Stdout.WriteString("Hello World!\n")
	fmt.Println(n, err)

	dirName := "dir"
	os.Mkdir(dirName, os.ModeDir)
	fileName := "file.txt"
	file, err := os.Create(fileName)
	file.WriteString("Hello World")
	file.Close()
	file, err = os.Open(fileName)
	buffer := make([]byte, 50)
	n, err = file.Read(buffer)
	fmt.Println(string(buffer))
	file.Close()
	os.Remove(dirName)
	os.Remove(fileName)

	fmt.Println(os.Getwd())
	fmt.Println(os.Hostname())
	fmt.Println(os.Environ()[13:15])
}
