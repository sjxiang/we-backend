package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

// 用户退出
func (h *UserHandler) Logout(c *gin.Context) {
	
	s := sessions.Default(c)
	s.Options(sessions.Options{
		MaxAge: -1,
	})
	// s.Clear()
	s.Save()
	
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "退出登录成功",
	})
}