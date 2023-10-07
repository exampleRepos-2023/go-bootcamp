package main

import (
	"context"
	"flag"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/masterclasses/microservices/solution/auth/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	serverAddr := flag.String(
		"addr",
		"localhost:50052",
		"The server address in the format of host:port")
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(*serverAddr, dialOptions...)
	if err != nil {
		log.Fatal("Unable to dial", err)
	}
	defer conn.Close()

	client := apis.NewAuthClient(conn)

	var username = uuid.New().String()
	var password = uuid.New().String()

	signUpResponse := signUp(client, username, password)
	signInResponse := signIn(client, username, password)
	signOutResponse := signOut(client, signInResponse)

	log.Println(signUpResponse.StatusCode, signUpResponse)
	log.Println(signInResponse.StatusCode, signInResponse)
	log.Println(signOutResponse.StatusCode, signOutResponse)
}

func signUp(authClient apis.AuthClient, username string, password string) *apis.SignUpResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	signUpRequest := apis.SignUpRequest{
		Username: username,
		Password: password,
	}
	signUpResponse, err := authClient.SignUp(ctx, &signUpRequest, grpc.EmptyCallOption{})
	if err != nil {
		log.Println("Unable to sign up", err)
		return nil
	}

	log.Println("Able to sign up", signUpResponse.StatusCode, signUpResponse)
	return signUpResponse
}

func signIn(authClient apis.AuthClient, username string, password string) *apis.SignInResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	signInRequest := apis.SignInRequest{
		Username: username,
		Password: password,
	}
	signInResponse, err := authClient.SignIn(ctx, &signInRequest, grpc.EmptyCallOption{})
	if err != nil {
		log.Println("Unable to sign in", err)
		return nil
	}

	log.Println("Able to sign in", signInResponse.StatusCode, signInResponse)
	return signInResponse
}

func signOut(authClient apis.AuthClient, authResponse *apis.SignInResponse) *apis.SignOutResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	signOutRequest := apis.SignOutRequest{
		SessionToken: authResponse.SessionToken,
	}
	signOutResponse, err := authClient.SignOut(ctx, &signOutRequest, grpc.EmptyCallOption{})
	if err != nil {
		log.Println("Unable to sign out", err)
		return nil
	}

	log.Println("Able to sign out", signOutResponse.StatusCode, signOutResponse)
	return signOutResponse
}
