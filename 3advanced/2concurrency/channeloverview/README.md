# Channels

## What is Channel
Channel is a Go data type, which provides means of communication between goroutines.

## Why do we need Channel?
```
func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start)) // Print out 1s
	}()
	go attack("Tommy")
	time.Sleep(time.Second)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at", target)
}
```

## Communication Models
Don't communicate by sharing memory; share memory by communicating.
- Communicate by sharing memory.
  
  Many languages share memory to communicate among concurrent processes. It's fully the developer's job to deal with synchronization, race condition when writing and reading of the shared memory.

- Share memory by communicating.

  Go uses channels to communicate messages between concurrent processes. The channel mechanism takes care of when the message is sent and received, and that each write and read is atomic.