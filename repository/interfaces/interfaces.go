package interfaces

import (
	"github.com/spacemono/go-api/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetAll() []*entity.User
	GetUserById(id string) *entity.User
}

type Hasher interface {
	Hash(string) string
}

type Validator interface {
	ValidateStruct(s any) error
}
