package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/spacemono/go-api/infrastructure/hash"
	"github.com/spacemono/go-api/infrastructure/repository"
	"github.com/spacemono/go-api/infrastructure/repository/postgres"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                                       // Allow all origins (change for production)
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allow common methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // Allowed headers
	)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", cors(router)))
	}()

	fmt.Println("Listening on port 8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	log.Println("Graceful shutdown...")

}
