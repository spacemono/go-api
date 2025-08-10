package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spacemono/go-api/domain/entity"
	"time"
)

func Generate(secretKey []byte, ttl time.Duration, user *entity.User) (string, error) {
	currTime := time.Now()
	expiresAtTime := currTime.Add(ttl)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(currTime),
		ExpiresAt: jwt.NewNumericDate(expiresAtTime),
		Subject:   user.Id,
	})

	return token.SignedString(secretKey)
}
