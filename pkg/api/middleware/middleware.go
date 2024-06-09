package middleware

import (
	"github.com/gin-gonic/gin"

	"we-backend/pkg/service/token"
)


type Middleware interface {
	Authenticate() gin.HandlerFunc
	EnableCORS() gin.HandlerFunc
}

type middleware struct {
	tokenService   token.TokenService
}

func NewMiddleware() Middleware {
	return &middleware{}
}

// 	ErrNoAuth                 = "请求头中的auth为空"
// 	ErrAuthFormatInvalid      = "请求头中的auth格式有错误"


