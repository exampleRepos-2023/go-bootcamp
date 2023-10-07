package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("New goroutine")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Main goroutine")

	// Mutex
	iterations := 1000
	sum := 0
	wg.Add(iterations)
	var mu sync.Mutex
	for i := 0; i < iterations; i++ {
		go func() {
			mu.Lock()
			sum++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(sum)

	// Once
	iterations = 1000
	sum = 0
	wg.Add(iterations)
	var once sync.Once
	for i := 0; i < iterations; i++ {
		go func() {
			once.Do(func() {
				sum++
			})
		}()
		wg.Done()
	}
	wg.Wait()
	fmt.Println(sum)

	// Pool
	memPool := &sync.Pool{
		New: func() interface{} {
			mem := make([]byte, 1024)
			return &mem
		},
	}
	mem := memPool.Get()
	// ...
	memPool.Put(mem)

	// Atomic
	var i int64
	atomic.AddInt64(&i, 1)
	mu.Lock()
	i += 1
	mu.Unlock()

	// Map
	syncMap := sync.Map{}
	iterations = 1000
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			syncMap.Store(0, i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Done with sync map")

}
