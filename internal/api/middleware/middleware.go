package middleware

import (
	"github.com/gin-gonic/gin"

	"we-backend/internal/service/accesscontrol"
	"we-backend/internal/service/token"
)


type Middleware interface {
	Authenticate() gin.HandlerFunc
	EnableCORS() gin.HandlerFunc
	RateLimit() gin.HandlerFunc 
}

type middleware struct {
	tokenService        token.TokenService
	rateLimitService    accesscontrol.AccessControl
}

func NewMiddleware(tokenService token.TokenService, rateLimitService accesscontrol.AccessControl) Middleware {
	return &middleware{
		tokenService:        tokenService,
		rateLimitService:    rateLimitService,
	}
}

// 	ErrNoAuth                 = "请求头中的auth为空"
// 	ErrAuthFormatInvalid      = "请求头中的auth格式有错误"


