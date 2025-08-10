package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/spacemono/go-api/domain/entity"
	"github.com/spacemono/go-api/domain/interfaces"
	"github.com/spacemono/go-api/service/command"
)

type User struct {
	userRepository interfaces.UserRepository
	hasher         interfaces.Hasher
	validator      *validator.Validate
}

func NewUser(userRepository interfaces.UserRepository, hasher interfaces.Hasher, validator *validator.Validate) *User {

	return &User{userRepository: userRepository, hasher: hasher, validator: validator}
}

func (s *User) GetAll() ([]*entity.User, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return nil, NewError(ErrBadRequest, err)
	}
	return users, nil
}

func (s *User) GetUserById(id string) (*entity.User, error) {
	user, err := s.userRepository.GetUserById(id)
	if err != nil {
		return nil, NewError(ErrBadRequest, errors.New("error in service"))
	}
	return user, nil
}

func (s *User) Create(cmd *command.CreateUser) (*command.CreateUserResult, error) {

	if err := s.validator.Struct(cmd); err != nil {
		return nil, NewError(ErrBadRequest, err)
	}

	user := &entity.User{Username: cmd.Username, Email: cmd.Email, PasswordHash: s.hasher.Hash(cmd.Password)}

	if err := s.userRepository.Create(user); err != nil {
		return nil, NewError(ErrBadRequest, errors.New("failed to create user in memory db"))
	}

	return &command.CreateUserResult{Username: cmd.Username, Email: cmd.Email, Password: cmd.Password}, nil
}
