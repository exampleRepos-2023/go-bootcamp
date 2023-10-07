package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Math struct{}

type Args struct {
	A, B int
}

type Reply struct {
	Sum int
}

func (t *Math) Add(args *Args, reply *Reply) error {
	reply.Sum = args.A + args.B
	return nil
}

func main() {
	// RPC service
	server := rpc.Server{}
	err := server.Register(new(Math))
	checkError(err)

	// Open the connection
	listener, err := net.Listen("tcp", ":1234")
	checkError(err)

	// Serve the listener
	protocol := "http" // tcp or http
	switch protocol {
	case "tcp":
		log.Println("RPC via TCP...")
		server.Accept(listener)
	case "http":
		log.Println("RPC via HTTP...")
		server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
		err = http.Serve(listener, &server)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
