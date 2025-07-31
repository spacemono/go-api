package postgres

import (
	"errors"
	"fmt"
	"github.com/spacemono/go-api/domain/entity"
	"github.com/spacemono/go-api/domain/interfaces"
	"log"
)

type UserRepository struct {
	client *Client
}

func NewUserRepository(client *Client) interfaces.UserRepository {

	return &UserRepository{client: client}
}

func (r *UserRepository) GetAll() ([]*entity.User, error) {
	var allUsers []*entity.User

	query := "SELECT * FROM users;"
	rows, err := r.client.db.Query(query)

	if err != nil {
		log.Fatal("No users in db")
	}

	for rows.Next() {
		var ID, username, name, passwordHash string

		if err := rows.Scan(&ID, &username, &name, &passwordHash); err != nil {
			return nil, errors.New("failed to get users from DB")
		}

		allUsers = append(allUsers, &entity.User{Id: ID, Username: username, Name: name, PasswordHash: passwordHash})
	}
	return allUsers, nil
}

func (r *UserRepository) GetUserById(id string) (*entity.User, error) {
	var ID, passwordHash, name, username string

	query := fmt.Sprintf("SELECT * FROM users WHERE id='%s';", id)

	row := r.client.db.QueryRow(query)
	err := row.Scan(&ID, &username, &name, &passwordHash)
	if err != nil {
		return nil, errors.New("failed to get user by this ID from DB")
	}

	return &entity.User{Id: ID, Username: username, Name: name, PasswordHash: passwordHash}, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	newUser := toDbUser(user)

	// RETURNING id

	query := fmt.Sprintf("INSERT INTO users (username, name, password_hash) VALUES ('%s', '%s', '%s')",
		newUser.Username, newUser.Name, newUser.PasswordHash)

	_, err := r.client.db.Query(query)
	if err != nil {
		return errors.New("failed to push user to DB")
	}

	return nil
}
