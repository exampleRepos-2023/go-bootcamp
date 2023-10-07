package main

import (
	"fmt"
	"gitlab.com/golangdojo/bootcamp/1beginner/4midtermproject1/solution/renderer"
)

func main() {
	err := renderer.R.Render()
	if err != nil {
		fmt.Println(err)
	}
}
