package middleware

import (
	"fmt"
	"net/http"
	"context"

	"github.com/gin-gonic/gin"
)


// 滑动窗口限流
func (h *middleware) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		if c.GetHeader("x-stress") == "true" {
			
			// 把压测标签 tag 带进 ctx
			newCtx := context.WithValue(c, StressKey, true)
			c.Request = c.Request.Clone(newCtx)
			c.Next()

			return
		}

		requestIP := requestIPKey(c.RemoteIP())

		over , err := h.accessControlService.Limit(c, requestIP)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)  // redis 被打崩了，那就别玩了
			return
		}

		if over {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"msg": "请求过载，请稍后重试"})
			return 
		}

		c.Next()
	}	
}


type ctxKey string

const (
	StressKey ctxKey = "x-stress"
)


func requestIPKey(ip string) string {
	return fmt.Sprintf("request_ip:%s", ip)
}