package interfaces

import (
	"github.com/spacemono/go-api/domain/entity"
	"github.com/spacemono/go-api/service/command"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetAll() ([]*entity.User, error)
	GetUserById(id string) (*entity.User, error)
}

type UserService interface {
	Create(user *entity.User) (*command.CreateUserResult, error)
	GetAll() ([]*entity.User, error)
	GetUserById(id string) (*entity.User, error)
}

type Hasher interface {
	Hash(string) string
}
