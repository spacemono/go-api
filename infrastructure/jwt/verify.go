package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidOrExpiredToken = errors.New("invalid or expired token")

func Verify(token string, secretKey []byte) (string, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if !t.Valid {
		return "", ErrInvalidOrExpiredToken
	}

	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", ErrInvalidOrExpiredToken
	}

	return claims.Subject, nil
}
