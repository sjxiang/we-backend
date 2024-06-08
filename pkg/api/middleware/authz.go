package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"we-backend/pkg/consts"
)

func (h *middleware) AuthenticateUserByCookie(ignorePaths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {	

		// 不需要校验的路由 
		for _, path := range ignorePaths {
			if c.Request.URL.Path == path {
				return
			}
		}

		// 校验登录
		value, err := c.Cookie("_cookie")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": true,
				"msg":   "需要登录",
			})
			return
		}


		userId, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": true,
				"msg":   "需要登录",
			})
			return
		}
		
		c.Set(consts.SessionKeyUserId, userId)
		c.Next()
	}
}
