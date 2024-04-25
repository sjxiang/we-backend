package middleware

import (
	"github.com/gin-gonic/gin"

	"we-backend/pkg/service/token"
)


type Middleware interface {
	HandleCors() gin.HandlerFunc 
	AuthenticateUser() gin.HandlerFunc
	AuthenticateUserByCookieAndSession() gin.HandlerFunc
}

type middleware struct {
	tokenService   token.TokenService

	// Todo. 如果了解 gin 注册中间件顺序的骚操作，就意识到多此一举
	ignorePaths    []string
}

func NewMiddleware() Middleware {
	return &middleware{
		ignorePaths: make([]string, 0),
	}
}

func (h *middleware) Build(paths ...string) Middleware {
	h.ignorePaths = paths
	return h
}


