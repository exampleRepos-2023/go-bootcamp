package credentials

import (
	"github.com/google/uuid"
	"testing"
)

func TestUsersImpl(t *testing.T) {
	U = NewUsers()
	username := uuid.New().String()
	password := uuid.New().String()

	err := U.CreateUser(username, password)
	if err != nil {
		t.Error("Unable to create user")
	}

	userUuid, exists := U.GetUserUuid(username, password)
	if !exists {
		t.Error("Unable to create user correctly")
	}

	err = U.CreateUser(username, password)
	if err == nil {
		t.Error("Unable to skip create the same user")
	}

	U.DeleteUser(userUuid)

	userUuid, exists = U.GetUserUuid(username, password)
	if exists {
		t.Error("Unable to delete user correctly")
	}
}