package main

import (
	"github.com/spacemono/go-api/infrastructure/hash"
	"log"
	"net/http"

	"github.com/spacemono/go-api/config"
	"github.com/spacemono/go-api/repository"
	"github.com/spacemono/go-api/repository/postgres"
	"github.com/spacemono/go-api/service"
	"github.com/spacemono/go-api/transport/rest"
)

func main() {
	cfg := config.New(".env")

	hasher := hash.NewHasher(cfg)

	db, err := postgres.NewClient(cfg)
	if err != nil {
		log.Fatal("Error: Failed connecting to DB", err)
	}

	repos := repository.New(db)
	userService := service.NewUser(repos.User, hasher)
	router := rest.NewHandler(userService)

	log.Fatal(http.ListenAndServe(":8080", router))
}
