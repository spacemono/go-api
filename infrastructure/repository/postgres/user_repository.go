package postgres

import (
	"errors"
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

	rows, err := r.client.db.Query("SELECT id, username, email, password_hash FROM users;")

	if err != nil {
		log.Fatal("No users in db")
	}

	for rows.Next() {
		var ID, username, email, passwordHash string

		if err := rows.Scan(&ID, &username, &email, &passwordHash); err != nil {
			return nil, errors.New("failed to get users from DB")
		}

		allUsers = append(allUsers, &entity.User{Id: ID, Username: username, Email: email, PasswordHash: passwordHash})
	}
	return allUsers, nil
}

func (r *UserRepository) GetUserById(id string) (*entity.User, error) {
	var ID, passwordHash, email, username string

	row := r.client.db.QueryRow("SELECT id, username, email, password_hash FROM users WHERE id=$1", id)
	err := row.Scan(&ID, &username, &email, &passwordHash)
	if err != nil {
		return nil, errors.New("failed to get user by this ID from DB")
	}

	return &entity.User{Id: ID, Username: username, Email: email, PasswordHash: passwordHash}, nil
}

func (r *UserRepository) Create(user *entity.User) error {

	// RETURNING id

	_, err := r.client.db.Query("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)",
		user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return errors.New("failed to push user to DB")
	}

	return nil
}
