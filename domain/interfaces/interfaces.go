package interfaces

import (
	"github.com/spacemono/go-api/domain/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetAll() ([]*entity.User, error)
	GetUserById(id string) (*entity.User, error)
}

type Hasher interface {
	Hash(string) string
}

type JwtManager interface {
	GenerateAccessToken(user *entity.User) (string, error)
	GenerateRefreshToken(user *entity.User) (string, error)
}
