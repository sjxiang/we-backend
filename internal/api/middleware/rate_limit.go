package middleware

import "github.com/gin-gonic/gin"

// 限流
func (h *middleware) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}