package main

import (
	"fmt"
	"time"
)

func main() {
	// Select statement lets a goroutine wait on multiple communication operations
	// It blocks the current goroutine until the soonest case is ready
	// or a random one if there are multiple
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "Message c1"
	}()
	go func() {
		c2 <- "Message c2"
	}()
	select {
	case m := <-c1:
		fmt.Println(m)
	case m := <-c2:
		fmt.Println(m)
	}
}
