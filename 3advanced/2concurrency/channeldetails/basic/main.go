package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	smokeSignal := make(chan string) // chan [message type]
	sender := (chan<- string)(smokeSignal)
	receiver := (<-chan string)(smokeSignal)
	go attack("Tommy", sender)
	go attack("Bobby", sender)
	fmt.Println(<-receiver)
	fmt.Println(<-receiver)
	// send >= receive
}

func attack(target string, attacked chan<- string) {
	fmt.Println("Throwing ninja stars at", target)
	attacked <- target
}
