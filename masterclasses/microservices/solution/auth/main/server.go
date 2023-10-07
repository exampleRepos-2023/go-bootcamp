package main

import (
	"gitlab.com/golangdojo/bootcamp/masterclasses/microservices/solution/auth/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	// register server
	server := apis.AuthService{}
	grpcServer := grpc.NewServer()
	apis.RegisterAuthServer(grpcServer, &server)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	// start to serve
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}