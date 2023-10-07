package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Package io provides basic interfaces to I/O(input/output) primitives
	// 1) Writer, Reader & Seeker interfaces
	// 2) Functions working alongside with these interfaces

	// Writer interface - wraps the basic Write method.
	// type Writer interface {
	//     Write(p []byte) (n int, err error)
	// }
	file, _ := os.Create("file.txt")
	writer := io.Writer(file)
	n, err := writer.Write([]byte("Hello"))
	fmt.Println(n, err)
	n, err = io.WriteString(writer, "!")
	fmt.Println(n, err)
	file.Close()

	// Reader interface -  wraps the basic Read method.
	// type Reader interface {
	//     Read(p []byte) (n int, err error)
	// }
	file, _ = os.Open("file.txt")
	reader := io.Reader(file)
	buffer := make([]byte, 10)
	n, err = reader.Read(buffer)
	fmt.Printf("Read n={%v}, err={%v}, buffer={%v}\n", n, err, string(buffer))
	file.Close()

	file, _ = os.Open("file.txt")
	reader = io.Reader(file)
	minBuffer := make([]byte, 1)
	for {
		n, err = reader.Read(minBuffer)
		fmt.Printf("%v, %v, %v\n", string(minBuffer), n, err)
		if err != nil {
			break
		}
	}
	file.Close()

	file, _ = os.Open("file.txt")
	reader = io.Reader(file)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("ReadAll buffer={%v}, err={%v}\n", string(buffer), err)
	file.Close()

	// Seeker interface - wraps the basic Seek method.
	// type Seeker interface {
	//     Seek(offset int64, whence int) (int64, error)
	// }
	file, _ = os.Open("file.txt")
	reader = io.Reader(file)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("ReadAll buffer={%v}, err={%v}\n", string(buffer), err)
	seeker := reader.(io.Seeker)
	seeker.Seek(-5, io.SeekCurrent)
	buffer, err = io.ReadAll(reader)
	fmt.Printf("ReadAll buffer={%v}, err={%v}\n", string(buffer), err)
	file.Close()
}
