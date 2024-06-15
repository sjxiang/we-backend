package middleware

import (
	"github.com/gin-gonic/gin"

	"we-backend/internal/service/access"
	"we-backend/internal/service/token"
)


type Middleware interface {
	Authenticate() gin.HandlerFunc
	EnableCORS() gin.HandlerFunc
	RateLimit() gin.HandlerFunc 
}

type middleware struct {
	tokenService         token.TokenService
	accessControlService access.AccessControlService
}

func NewMiddleware(tokenService token.TokenService, accessControlService access.AccessControlService) Middleware {
	return &middleware{
		tokenService:         tokenService,
		accessControlService: accessControlService,
	}
}

// 	ErrNoAuth                 = "请求头中的auth为空"
// 	ErrAuthFormatInvalid      = "请求头中的auth格式有错误"


