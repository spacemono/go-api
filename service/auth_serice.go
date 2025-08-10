package service

import "github.com/spacemono/go-api/domain/interfaces"

type AuthService struct {
	jwtManager interfaces.JwtManager
}

func NewAuthService(jwtManager interfaces.JwtManager) *AuthService {
	return &AuthService{jwtManager: jwtManager}
}
