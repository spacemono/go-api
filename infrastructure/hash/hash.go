package hash

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/spacemono/go-api/config"
)

type Hasher struct {
	salt string
}

func NewHasher(cfg *config.Config) *Hasher {

	return &Hasher{salt: cfg.Hasher.HashSalt}
}

func (h *Hasher) Hash(txt string) string {
	hash := sha256.New()
	hash.Write([]byte(h.salt + txt))
	return hex.EncodeToString(hash.Sum(nil))
}
