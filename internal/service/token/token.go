package token

import (
	"time"

	"we-backend/internal/conf"
)

type TokenService interface {
	CreateToken(id int64, email string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

func NewTokenService(cfg *conf.Config) TokenService {
	return &JWTMaker{cfg.SecretKey}
}
