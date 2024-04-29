package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 退出
func (h *AuthHandler) AuthzLogoutByCookie(c *gin.Context) {
	
	c.SetCookie(
		DEFAULT_COOKIE_NAME,
		"",
		-1, 
		DEFAULT_COOKIE_PATH, 
		DEFAULT_COOKIE_DOMAIN, 
		DEFAULT_COOKIE_SECURE, 
		DEFAULT_COOKIE_HTTP_ONLY,
	)

	// feedback
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "退出当前会话",
	})
}