package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spacemono/go-api/config"
)

type Client struct {
	db *sql.DB
}

func NewClient(cfg *config.Config) (*Client, error) {

	db, err := sql.Open("postgres", cfg.Postgres.Url)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	return &Client{db: db}, nil
}
