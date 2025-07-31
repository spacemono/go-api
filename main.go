package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/spacemono/go-api/infrastructure/hash"
	"github.com/spacemono/go-api/infrastructure/repository"
	"github.com/spacemono/go-api/infrastructure/repository/postgres"
	"log"
	"net/http"

	"github.com/spacemono/go-api/config"
	"github.com/spacemono/go-api/service"
	"github.com/spacemono/go-api/transport/rest"
)

func main() {
	cfg := config.New(".env")

	hasher := hash.NewHasher(cfg)
	validate := validator.New()

	db, err := postgres.NewClient(cfg)
	if err != nil {
		log.Fatal("Error: Failed connecting to DB", err)
	}

	repos := repository.New(db)
	userService := service.NewUser(repos.User, hasher, validate)
	router := rest.NewHandler(userService)

	log.Fatal(http.ListenAndServe(":8080", router))
}
