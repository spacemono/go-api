package interfaces

import (
	"github.com/spacemono/go-api/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetAll() ([]*entity.User, error)
	GetUserById(id string) (*entity.User, error)
}

type UserService interface {
}

type Hasher interface {
	Hash(string) string
}

type Validator interface {
	ValidateStruct(s any) error
}
