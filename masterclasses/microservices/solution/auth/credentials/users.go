package credentials

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var U = NewUsers()

type Users interface {
	CreateUser(username string, password string) error
	GetUserUuid(username string, password string) (uuid.UUID, bool)
	DeleteUser(userUuid uuid.UUID)
}

type UsersImpl struct {
	userUuidToUser map[uuid.UUID]User
	usernameToUser map[string]User
}

type User struct {
	userUuid uuid.UUID
	username string
	password string
}

func NewUsers() Users {
	return &UsersImpl{
		userUuidToUser: map[uuid.UUID]User{},
		usernameToUser: map[string]User{},
	}
}

func (u *UsersImpl) CreateUser(username string, password string) error {
	if _, exists := u.usernameToUser[username]; exists {
		return errors.New("unable to create an existing user")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Unable to encrypt password", err)
		return err
	}

	user := User{
		userUuid: uuid.New(),
		username: username,
		password: string(encryptedPassword),
	}

	u.usernameToUser[user.username] = user
	u.userUuidToUser[user.userUuid] = user

	return nil
}

func (u *UsersImpl) GetUserUuid(username string, password string) (uuid.UUID, bool) {
	user, exists := u.usernameToUser[username]
	if !exists {
		return uuid.UUID{}, false
	}

	if user.username == username &&
		bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password)) == nil {
		return user.userUuid, true
	}

	return uuid.UUID{}, false
}

func (u *UsersImpl) DeleteUser(userUuid uuid.UUID) {
	user, exists := u.userUuidToUser[userUuid]
	if exists {
		delete(u.userUuidToUser, userUuid)
		delete(u.usernameToUser, user.username)
	}
}
