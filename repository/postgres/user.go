package postgres

import (
	"github.com/spacemono/go-api/entity"
)

type User struct {
	Id           string
	Username     string
	Name         string
	PasswordHash string
}

func toDbUser(user *entity.User) *User {

	return &User{Username: user.Username, Name: user.Name, PasswordHash: user.PasswordHash}
}

func fromDbUser(user *User) *entity.User {
	return &entity.User{Username: user.Username, Name: user.Name, PasswordHash: user.PasswordHash}
}
