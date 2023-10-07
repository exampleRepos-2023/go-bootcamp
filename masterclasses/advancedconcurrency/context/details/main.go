package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	var handler http.ServeMux
	handler.HandleFunc("/", handleRequest)
	server := http.Server{
		Addr:         "",
		Handler:      &handler,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	var ctx    context.Context
	var cancel context.CancelFunc

	overallTimeout, err := time.ParseDuration(req.FormValue("overallTimeout"))
	perCallTimeout, err := time.ParseDuration(req.FormValue("perCallTimeout"))
	ctx = context.Background()
	//ctx = context.TODO()
	if err == nil {
		ctx, cancel = context.WithTimeout(ctx, perCallTimeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	go func() {
		time.Sleep(overallTimeout)
		cancel()
	}()

	anotherServerCall(ctx)

	<-ctx.Done()
	err = ctx.Err()
	if err != nil {
		return
	}
}

func anotherServerCall(ctx context.Context) {
	log.Println(ctx.Deadline())

	ctx = context.WithValue(ctx, "requestId", "123")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Before anotherServerCall2")
	anotherServerCall2(ctx)
	log.Println("After anotherServerCall2")
}

func anotherServerCall2(ctx context.Context) {
	log.Println(ctx.Value("requestId"))
	for {
		deadline, ok := ctx.Deadline()
		if !ok {
			break
		}
		if deadline.Before(time.Now()) {
			break
		}
		time.Sleep(1*time.Second)
	}
}
