package repository

import (
	"github.com/spacemono/go-api/domain/interfaces"
	postgres2 "github.com/spacemono/go-api/infrastructure/repository/postgres"
)

type Repositories struct {
	User interfaces.UserRepository
}

func New(client *postgres2.Client) *Repositories {
	user := postgres2.NewUserRepository(client)

	return &Repositories{User: user}

}
