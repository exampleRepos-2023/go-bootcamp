package main

import (
	"fmt"
	"time"
)

func main() {
	// Package time provides functionality for measuring and displaying time
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Date())
	fmt.Println(t.Month())
	fmt.Println(t.UTC())

	fmt.Println(t.Weekday() == time.Friday)
	fmt.Println(t.Location() == time.UTC)

	t = time.Date(2050, 1, 1, 1, 1, 1, 1, time.UTC)
	fmt.Println(t)

	//time.Sleep(time.Second * 2)
	t = time.Now()
	time.Sleep(time.Until(time.Now().Add(time.Second * 2)))
	fmt.Println(time.Since(t))

	fmt.Println("Channel Related:")
	fmt.Println(time.Now())
	ch := time.After(time.Second)
	fmt.Println(<-ch)
}
