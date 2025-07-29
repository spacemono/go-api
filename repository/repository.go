package repository

import (
	"github.com/spacemono/go-api/repository/interfaces"
	"github.com/spacemono/go-api/repository/postgres"
)

type Repositories struct {
	User interfaces.UserRepository
}

func New(client *postgres.Client) *Repositories {
	user := postgres.NewUserRepository(client)

	return &Repositories{User: user}
}
