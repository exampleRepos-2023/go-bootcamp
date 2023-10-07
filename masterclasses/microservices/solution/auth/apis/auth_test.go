package apis

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gitlab.com/golangdojo/bootcamp/masterclasses/microservices/solution/auth/credentials"
	"testing"
	"time"
)

type UsersMock struct {
	CreateUserError error
	CreateUserCalled bool
	GetUserUserUuid uuid.UUID
	GetUserExists bool
	GetUserCalled bool
	DeleteUserCalled bool
}

func (u *UsersMock) CreateUser(username string, password string) error {
	u.CreateUserCalled = true
	return u.CreateUserError
}

func (u *UsersMock) GetUserUuid(username string, password string) (uuid.UUID, bool) {
	u.GetUserCalled = true
	return u.GetUserUserUuid, u.GetUserExists
}

func (u *UsersMock) DeleteUser(userUuid uuid.UUID) {
	u.DeleteUserCalled = true
}

type SessionsMock struct {
	CreateSessionSessionUuid uuid.UUID
	CreateSessionCalled bool
	CheckSessionExists bool
	CheckSessionCalled bool
	DeleteSessionCalled bool
}

func (s *SessionsMock) CreateSession(userUuid uuid.UUID) uuid.UUID {
	s.CreateSessionCalled = true
	return s.CreateSessionSessionUuid
}

func (s *SessionsMock) CheckSession(token uuid.UUID) bool {
	s.CheckSessionCalled = true
	return s.CheckSessionExists
}

func (s *SessionsMock) DeleteSession(token uuid.UUID) {
	s.DeleteSessionCalled = true
}

func TestSignUp_Success(t *testing.T) {
	usersMock := UsersMock{
		CreateUserError: nil,
		CreateUserCalled: false,
	}
	credentials.U = &usersMock

	authService := AuthService{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := SignUpRequest{
		Username: uuid.New().String(),
		Password: uuid.New().String(),
	}

	response, err := authService.SignUp(ctx, &request)
	if err != nil {
		t.Error("Unable to sign up", err)
	}
	if response.StatusCode != StatusCode_SUCCESS {
		t.Error("Unable to sign up successfully")
	}
	if !usersMock.CreateUserCalled {
		t.Error("Unable to call create user")
	}
}

func TestSignUp_Failure(t *testing.T) {
	usersMock := UsersMock{
		CreateUserError: errors.New("unable to create user"),
		CreateUserCalled: false,
	}
	credentials.U = &usersMock

	authService := AuthService{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := SignUpRequest{
		Username: uuid.New().String(),
		Password: uuid.New().String(),
	}

	response, err := authService.SignUp(ctx, &request)
	if err == nil {
		t.Error("Unable to get sign up error", err)
	}
	if response.StatusCode != StatusCode_FAILURE {
		t.Error("Unable to get sign up failure")
	}
	if !usersMock.CreateUserCalled {
		t.Error("Unable to call create user")
	}
}

func TestSignIn_Success(t *testing.T) {
	userUuid := uuid.New()
	usersMock := UsersMock{
		GetUserUserUuid: userUuid,
		GetUserExists: true,
		GetUserCalled: false,
	}
	sessionToken := uuid.New()
	sessionsMock := SessionsMock{
		CreateSessionSessionUuid: sessionToken,
		CreateSessionCalled: false,
	}
	credentials.U = &usersMock
	credentials.S = &sessionsMock

	authService := AuthService{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := SignInRequest{
		Username: uuid.New().String(),
		Password: uuid.New().String(),
	}

	response, err := authService.SignIn(ctx, &request)
	if err != nil {
		t.Error("Unable to sign in", err)
	}
	if response.StatusCode != StatusCode_SUCCESS {
		t.Error("Unable to sign in successfully")
	}
	if !usersMock.GetUserCalled {
		t.Error("Unable to call get user")
	}
	if !sessionsMock.CreateSessionCalled {
		t.Error("Unable to call create session")
	}
}

func TestSignIn_Failure(t *testing.T) {
	userUuid := uuid.New()
	usersMock := UsersMock{
		GetUserUserUuid: userUuid,
		GetUserExists: false,
		GetUserCalled: false,
	}
	sessionToken := uuid.New()
	sessionsMock := SessionsMock{
		CreateSessionSessionUuid: sessionToken,
		CreateSessionCalled: false,
	}
	credentials.U = &usersMock
	credentials.S = &sessionsMock

	authService := AuthService{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := SignInRequest{
		Username: uuid.New().String(),
		Password: uuid.New().String(),
	}

	response, err := authService.SignIn(ctx, &request)
	if err == nil {
		t.Error("Unable to get sign in error", err)
	}
	if response.StatusCode != StatusCode_FAILURE {
		t.Error("Unable to get sign in failure")
	}
	if !usersMock.GetUserCalled {
		t.Error("Unable to call get user")
	}
	if sessionsMock.CreateSessionCalled {
		t.Error("Unable to skip create session")
	}
}

func TestSignOut_Success(t *testing.T) {
	sessionsMock := SessionsMock{
		DeleteSessionCalled: false,
	}
	credentials.S = &sessionsMock

	authService := AuthService{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := SignOutRequest{
		SessionToken: uuid.New().String(),
	}

	response, err := authService.SignOut(ctx, &request)
	if err != nil {
		t.Error("Unable to sign out", err)
	}
	if response.StatusCode != StatusCode_SUCCESS {
		t.Error("Unable to sign out successfully")
	}
	if !sessionsMock.DeleteSessionCalled {
		t.Error("Unable to call delete session")
	}
}

func TestSignOut_Failure(t *testing.T) {
	sessionsMock := SessionsMock{
		DeleteSessionCalled: false,
	}
	credentials.S = &sessionsMock

	authService := AuthService{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := SignOutRequest{
		SessionToken: "Invalid UUID string",
	}

	response, err := authService.SignOut(ctx, &request)
	if err == nil {
		t.Error("Unable to get sign out error", err)
	}
	if response.StatusCode != StatusCode_FAILURE {
		t.Error("Unable to get sign out failure")
	}
	if sessionsMock.DeleteSessionCalled {
		t.Error("Unable to skip delete session")
	}
}