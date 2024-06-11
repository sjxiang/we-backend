package token

import (
	"time"

	"we-backend/pkg/config"
)

type TokenService interface {
	CreateToken(id int64, email string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

func NewTokenService(cfg *config.Config) TokenService {
	return &JWTMaker{cfg.SecretKey}
}
