package postgres

import (
	"errors"
	"fmt"
	"log"

	"github.com/spacemono/go-api/entity"
	"github.com/spacemono/go-api/repository/interfaces"
)

type UserRepository struct {
	client *Client
}

func NewUserRepository(client *Client) interfaces.UserRepository {

	return &UserRepository{client: client}
}

func (r *UserRepository) GetAll() []*entity.User {
	allUsers := []*entity.User{}

	query := "SELECT * FROM users;"
	rows, err := r.client.db.Query(query)

	if err != nil {
		log.Fatal("No users in db")
	}

	for rows.Next() {
		var ID, username, name, passwordHash string
		if err := rows.Scan(&ID, &username, &name, &passwordHash); err != nil {
			log.Fatal("Error getting data")
		}
		allUsers = append(allUsers, &entity.User{Id: ID, Username: username, Name: name, PasswordHash: passwordHash})
	}
	return allUsers
}

func (r *UserRepository) GetUserById(id string) *entity.User {
	var ID, passwordHash, name, username string

	query := fmt.Sprintf("SELECT * FROM users WHERE id='%s';", id)

	row := r.client.db.QueryRow(query)
	err := row.Scan(&ID, &username, &name, &passwordHash)
	if err != nil {
		log.Fatal("Can't scan the user: ", err)
	}

	return &entity.User{Id: ID, Username: username, Name: name, PasswordHash: passwordHash}
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
