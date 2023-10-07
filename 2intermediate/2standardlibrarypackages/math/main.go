package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Package math
	// Mathematical constants and functions
	fmt.Println(math.Pi)
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt32)

	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1.31, 2))
	fmt.Println(math.Ceil(1.5))
	fmt.Println(math.Floor(1.5))
	fmt.Println(math.Round(1.49))
	fmt.Println(math.Round(1.50))
	fmt.Println(math.Pow(2, 10))
	fmt.Println(math.Sqrt(9))
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Int())
	fmt.Println(rand.Float64())
}
