package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres Postgres
	Hasher   Hasher
}

type Postgres struct {
	Url string `env:"POSTGRES_URL"`
}

type Hasher struct {
	HashSalt string `env:"HASH_SALT"`
}

func New(fileNames ...string) *Config {
	if err := godotenv.Load(fileNames...); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", fileNames, err)
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatal("Error creating config")
	}

	return cfg
}
