package main

import (
	"gitlab.com/golangdojo/bootcamp/3advanced/3webdevelopment/rpcclient/model"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Math struct{}

func (t *Math) Add(args *model.Args, reply *model.Reply) error {
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
	protocol := model.HTTP // tcp or http
	switch protocol {
	case model.TCP:
		log.Println("RPC via TCP...")
		server.Accept(listener)
	case model.HTTP:
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
