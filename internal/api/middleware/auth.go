package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey       = "authorization"
	authorizationTypeBearer      = "bearer"
	authorizationPayloadIDKey    = "authorization_payload_id"
	authorizationPayloadEmailKey = "authorization_payload_email"
)


// Authenticate anthn 认证
func (h *middleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("authorization header is not provided"))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("invalid authorization header format"))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("unsupported authorization type %s", authorizationType))
			return
		}

		// 校验 token
		accessToken := fields[1]
		payload, err := h.tokenService.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		// 会话保持
		if time.Until(payload.ExpiredAt) < time.Minute * time.Duration(30)  {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("时间差不多喽"))  // 考虑再续一轮
			return
		}

		c.Set(authorizationPayloadIDKey, payload.ID)
		c.Set(authorizationPayloadEmailKey, payload.Email)
		
		c.Next()
	}
}


// Authorize authz 授权
func (mw middleware) Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
	}
}