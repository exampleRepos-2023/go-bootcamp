package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	Pipeline:

	    p1          t1              a1
	d1--p2--d2--(ch)t2--d3--(ch)a2--d1
	    p3          t3              a3
	 */

	problems := []int{1, 2, 3}

	tasks := source(problems)

	achievement1 := perform(tasks)

	consume(achievement1)

	// This isn't quite parallel yet
	// but sets things up for more complicated patterns
}

func source(problems []int) <-chan int {
	tasks := make(chan int)
	go func() {
		for _, task := range problems {
			tasks <- task
		}
		close(tasks)
	}()
	return tasks
}

func perform(tasks <-chan int) <-chan int {
	achievements := make(chan int)
	go func() {
		for piece := range tasks {
			result := piece * piece
			time.Sleep(1*time.Second)
			achievements<- result
		}
		close(achievements)
	}()
	return achievements
}

func consume(achievement <-chan int) {
	for solution := range achievement {
		fmt.Println(solution)
	}
}
