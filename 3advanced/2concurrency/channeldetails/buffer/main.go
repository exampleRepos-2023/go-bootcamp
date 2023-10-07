package main

import "fmt"

func main() {
	// Buffered channel
	smokeSignal := make(chan string, 2)
	// smokeSignal := make(chan string, 0)
	smokeSignal <- "Tommy"
	smokeSignal <- "Bobby"
	fmt.Println(<-smokeSignal)
	fmt.Println(<-smokeSignal)
	/*
			main goroutine            smokeSignal
			                               -
		        "Tommy"                |      |
		                          |       ||       |
	*/
}
