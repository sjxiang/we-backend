package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


// 限流
func (h *middleware) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		over , err := h.rateLimitService.Limit(c, fmt.Sprintf("request_ip:%s", c.RemoteIP()))
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)  // redis 被打崩了
			return
		}

		if !over {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"msg": "请求过载，请稍后重试"})
			return 
		}

		c.Next()
	}	
}

