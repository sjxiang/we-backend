package middleware

import (
	"github.com/gin-gonic/gin"

	"we-backend/pkg/service/token"
)


type Middleware interface {
	Cors() gin.HandlerFunc 
	AuthenticateUserByJWT() gin.HandlerFunc

	AuthenticateUserByCookie(ignorePaths ...string) gin.HandlerFunc
	AuthenticateUserByCookieAndSession(ignorePaths ...string) gin.HandlerFunc 
}

type middleware struct {
	tokenService   token.TokenService
}

func NewMiddleware() Middleware {
	return &middleware{}
}


