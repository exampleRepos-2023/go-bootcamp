package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutine
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	evilNinjas := []string{"Tommy", "Bobby", "Andy"}
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}
	time.Sleep(time.Second * 2)
}

func attack(target string) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
}
