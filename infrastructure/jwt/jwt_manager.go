package jwt

import (
	"github.com/spacemono/go-api/config"
	"github.com/spacemono/go-api/domain/entity"
	"github.com/spacemono/go-api/domain/interfaces"
)

type Manager struct {
	cfg *config.JWT
}

func NewManager(cfg *config.JWT) interfaces.JwtManager {
	return &Manager{cfg: cfg}
}

func (m *Manager) GenerateAccessToken(user *entity.User) (string, error) {
	return Generate(m.cfg.AccessSecret, m.cfg.AccessTTL, user)
}

func (m *Manager) GenerateRefreshToken(user *entity.User) (string, error) {
	return Generate(m.cfg.RefreshSecret, m.cfg.RefreshTTL, user)
}

func (m *Manager) VerifyAccessToken(token string) (string, error) {
	return Verify(token, m.cfg.AccessSecret)
}

func (m *Manager) VerifyRefreshToken(token string) (string, error) {
	return Verify(token, m.cfg.RefreshSecret)
}
