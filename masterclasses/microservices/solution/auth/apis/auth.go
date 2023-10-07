package apis

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/masterclasses/microservices/solution/auth/credentials"
	"log"
)

type AuthService struct {

}

func (a *AuthService) SignUp(_ context.Context, request *SignUpRequest) (*SignUpResponse, error) {
	err := credentials.U.CreateUser(request.Username, request.Password)

	if err != nil {
		log.Println("Unable to create user", err)
		return &SignUpResponse{
			StatusCode: StatusCode_FAILURE,
		}, errors.New("unable to create user")
	}

	log.Println("Able to sign up")

	return &SignUpResponse{
		StatusCode: StatusCode_SUCCESS,
	}, nil
}

func (a *AuthService) SignIn(_ context.Context, request *SignInRequest) (*SignInResponse, error) {
	userUuid, exists := credentials.U.GetUserUuid(request.Username, request.Password)
	if !exists {
		log.Println("Unable to get user UUID")
		return &SignInResponse{
			StatusCode: StatusCode_FAILURE,
		}, errors.New("unable to get UserUuid")
	}

	sessionToken := credentials.S.CreateSession(userUuid)

	log.Println("Able to sign in")

	return &SignInResponse{
		StatusCode:   StatusCode_SUCCESS,
		SessionToken: sessionToken.String(),
		UserUuid:     userUuid.String(),
	}, nil
}

func (a *AuthService) SignOut(_ context.Context, request *SignOutRequest) (*SignOutResponse, error) {
	sessionToken, err := uuid.Parse(request.SessionToken)
	if err != nil {
		log.Println("Unable to parse request session token")
		return &SignOutResponse{
			StatusCode: StatusCode_FAILURE,
		}, err
	}

	credentials.S.DeleteSession(sessionToken)
	log.Println("Able to sign out")
	return &SignOutResponse{
		StatusCode: StatusCode_SUCCESS,
	}, nil
}

func (a *AuthService) mustEmbedUnimplementedAuthServer() {

}