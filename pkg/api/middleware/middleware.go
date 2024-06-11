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

func NewMiddleware(tokenService token.TokenService) Middleware {
	return &middleware{tokenService: tokenService}
}

// 	ErrNoAuth                 = "请求头中的auth为空"
// 	ErrAuthFormatInvalid      = "请求头中的auth格式有错误"


