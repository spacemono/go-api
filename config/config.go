package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres Postgres
	Hasher   Hasher
	JWT      JWT
}

type Postgres struct {
	Url string `env:"POSTGRES_URL"`
}

type Hasher struct {
	HashSalt string `env:"HASH_SALT"`
}

type JWT struct {
	AccessSecret  string        `env:"JWT_ACCESS_SECRET"`
	RefreshSecret string        `env:"JWT_REFRESH_SECRET"`
	AccessTTL     time.Duration `env:"JWT_ACCESS_TTL"`
	RefreshTTL    time.Duration `env:"JWT_REFRESH_TTL"`
}

func New(fileNames ...string) *Config {
	if err := godotenv.Load(fileNames...); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", fileNames, err)
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Error creating config %v", err)
	}

	return cfg
}
