package middleware

import (
	"time"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

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
		if err != nil || value == "" {
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



func (h *middleware) AuthenticateUserByCookieAndSession(ignorePaths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {	

		// 不需要校验的路由 
		for _, path := range ignorePaths {
			if c.Request.URL.Path == path {
				return
			}
		}

		// 校验登录
		session := sessions.Default(c)
		
		uid, last := session.Get(consts.SessionKeyUserId), session.Get(consts.SessionKeyLastTime)
		if uid == nil || last == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": true,
				"msg":   "需要登录",
			})
			return
		}
	
		// 会话状态保持
		if lastTime, ok := last.(time.Time); ok {
			if time.Since(lastTime) > time.Minute*30 {  // 还有 30 min，刷新，再给你续一轮
				session.Options(sessions.Options{
					MaxAge: 7 * 86400,
				})
				session.Save()
				
				log.Info().Msg("续了一轮")
			} 	
		}
		

		c.Set(consts.SessionKeyUserId, uid.(int64))
		c.Next()
	}
}


