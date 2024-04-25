package middleware

import (
	"net/http"
	"time"
	"we-backend/pkg/consts"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *middleware) AuthenticateUserByCookieAndSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		// 不需要校验的路由 
		for _, path := range h.ignorePaths {
			if c.Request.URL.Path == path {
				return
			}
		}

		session := sessions.Default(c) 
		userId := session.Get(consts.SessionKeyUserId)
		lastTime := session.Get(consts.SessionKeyLastTime)

		if userId == nil || lastTime == nil {
			// 没有登录
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": true,
				"msg":   "需要登录",
			})
			return
		}

		// 会话状态保持
		if last, ok := lastTime.(time.Time); ok {
			if time.Since(last) > time.Minute*30 {  // 还有 30 min，刷新，再给你续一轮
				session.Options(sessions.Options{
					MaxAge: 7 * 86400,
				})
				session.Save()
				// zap.L().Info("续了一轮")
			} 	
		}

		c.Set(consts.SessionKeyUserId, userId.(int64))
		c.Next()

	}
}
