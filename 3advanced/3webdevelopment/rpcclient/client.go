package main

import (
	"fmt"
	"gitlab.com/golangdojo/bootcamp/3advanced/3webdevelopment/rpcclient/model"
	"math/rand"
	"net/rpc"
	"time"
)

func main() {
	// RPC main

	// Connect to an RPC handlers
	var client *rpc.Client
	protocol := model.HTTP // tcp or http
	switch protocol {
	case model.TCP:
		client, _ = rpc.Dial("tcp", ":1234")
	case model.HTTP:
		client, _ = rpc.DialHTTP("tcp", ":1234")
	}

	// Synchronous call
	rand.Seed(time.Now().Unix())
	a, b := rand.Int()%11, rand.Int()%11
	args := model.Args{A: a, B: b}
	var reply model.Reply
	err := client.Call("Math.Add", &args, &reply)
	fmt.Printf("Math.Add: %d+%d=%d with error=%v\n",
		args.A, args.B, reply.Sum, err)

	// Asynchronous call
	reply = model.Reply{}
	asyncCall := client.Go("Math.Add", &args, &reply, nil)
	<-asyncCall.Done
	fmt.Printf("Math.Add: %d+%d=%d with error=%v\n",
		args.A, args.B, reply.Sum, asyncCall.Error)
}
