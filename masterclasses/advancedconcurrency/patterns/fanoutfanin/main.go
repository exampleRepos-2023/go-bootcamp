package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
	Pipeline:

		    p1          t1          a1
		d1--p2--d2--(ch)t2--d3--(ch)a2--d1
		    p3          t3          a3

	Fan-out fan-in:

			p1          t1  w1--(ch)a1          a1
		d1--p2--d2--(ch)t2--w2--(ch)a2--d1--(ch)a2--d1
			p3          t3  w3--(ch)a3          a3

			p1          t1  w1--(ch)a1--w4--(ch)a1          a1
		d1--p2--d2--(ch)t2--w2--(ch)a2--w5--(ch)a2--d1--(ch)a2--d1
			p3          t3  w3--(ch)a3--w6--(ch)a3          a3
	*/
	problems := []int{1, 2, 3}

	tasks := source(problems)

	//achievement1 := perform(tasks)
	//achievement2 := perform(tasks)
	//
	//consume(achievement1)
	//consume(achievement2)

	var achievements []<-chan int
	workerNum := 1
	for workerI := 0; workerI < workerNum; workerI++ {
		achievement1 := perform(tasks)
		achievement2 := perform(achievement1)
		achievements = append(achievements, achievement1, achievement2)
	}

	combinedAchievement := merge(achievements)
	consume(combinedAchievement)
}

func source(problems []int) <-chan int {
	ninjaTasks := make(chan int)
	go func() {
		for _, ninjaTask := range problems {
			ninjaTasks<- ninjaTask
		}
		close(ninjaTasks)
	}()
	return ninjaTasks
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

func merge(achievements []<-chan int) <-chan int {
	var wg sync.WaitGroup
	combinedAchievement := make(chan int)

	output := func(achievement <-chan int) {
		for solution := range achievement {
			combinedAchievement <- solution
		}
		wg.Done()
	}

	wg.Add(len(achievements))
	for _, achievement := range achievements {
		go output(achievement)
	}

	go func() {
		wg.Wait()
		close(combinedAchievement)
	}()
	return combinedAchievement
}
