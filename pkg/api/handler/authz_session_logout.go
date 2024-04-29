package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 退出
func (h *AuthHandler) AuthzLogoutBySession(c *gin.Context) {
	
	session := sessions.Default(c)

	session.Options(sessions.Options{
		MaxAge: -1,
	})
	session.Save()
	
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "退出登录",
	})
	
}