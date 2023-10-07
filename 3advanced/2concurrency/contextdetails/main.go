package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Context is a package that communicates
	// values, deadlines, cancellation signals between processes

	// Creating a blank context without any values or deadlines
	ctx := context.TODO()
	ctx = context.Background()

	// Value
	key := "request-id"
	value := "123"
	ctxValue := context.WithValue(ctx, key, value)
	fmt.Println(ctx.Value(key))
	fmt.Println(ctxValue.Value(key))

	childKey := "child-request-id"
	childValue := "456"
	childCtxValue := context.WithValue(ctxValue, childKey, childValue)
	fmt.Println(childCtxValue.Value(key))
	fmt.Println(childCtxValue.Value(childKey))
	fmt.Println(ctxValue.Value(childKey))

	go func(ctx context.Context) {
		fmt.Println("Reading context from different goroutine:", childCtxValue.Value(childKey))
	}(ctx)

	// Same for the deadlines and cancellation signals we will set next

	// Deadlines
	ctxDeadline, _ := context.WithDeadline(ctx, time.Now().Add(time.Second))
	//ctxTimeout, cancelTimeout := context.WithTimeout(ctx, time.Second)
	//ctxDeadline, cancelDeadline := context.WithDeadline(ctx, time.Now().Add(time.Second))
	//defer cancelDeadline()
	//cancelDeadline() // Context is done when cancelled

	select {
	case <-ctxDeadline.Done():
		fmt.Println("Context reached here first: ", ctxDeadline.Err())
	case <-time.After(time.Second * 2):
		fmt.Println("Process reached here first:", ctxDeadline.Err())
	}

	// Cancellation Signals
	//ctxDeadline, cancelDeadline := context.WithDeadline(ctx, time.Now().Add(time.Second))
	//defer cancelDeadline()
	//cancelDeadline() // Context is done when cancelled
}
