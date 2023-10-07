package main

import (
	"fmt"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/problem/stage2/renderer"
)

func main() {
	err := renderer.R.Render()
	if err != nil {
		fmt.Println(err)
	}
}
