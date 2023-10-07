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
	evilNinjas := []string{"Tommy", "Johnny"}
	smokeSignal := make(chan string)
	go attack(evilNinjas, smokeSignal)
	for message := range smokeSignal {
		fmt.Println(message)
	}
}

func attack(evilNinjas []string, attacked chan string) {
	for _, evilNinja := range evilNinjas {
		fmt.Println("Throwing ninja stars at", evilNinja)
		attacked <- evilNinja
	}
	close(attacked)
}
