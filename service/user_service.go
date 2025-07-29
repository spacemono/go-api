package service

import (
	"errors"

	"github.com/spacemono/go-api/entity"
	"github.com/spacemono/go-api/repository/interfaces"
	"github.com/spacemono/go-api/service/command"
)

type User struct {
	userRepository interfaces.UserRepository
	hasher         interfaces.Hasher
}

func NewUser(userRepository interfaces.UserRepository, hasher interfaces.Hasher) *User {

	return &User{userRepository: userRepository, hasher: hasher}
}

func (s *User) GetAll() []*entity.User {
	users := s.userRepository.GetAll()
	return users
}

func (s *User) GetUserById(id string) *entity.User {
	user := s.userRepository.GetUserById(id)
	return user
}

func (s *User) Create(cmd *command.CreateUser) (*command.CreateUserResult, error) {
	user := &entity.User{Name: cmd.Name, Username: cmd.Username, PasswordHash: s.hasher.Hash(cmd.Password)}

	if err := s.userRepository.Create(user); err != nil {
		return nil, errors.New("user can not be created")
	}

	return &command.CreateUserResult{Username: cmd.Username, Name: cmd.Name, Password: cmd.Password}, nil
}
