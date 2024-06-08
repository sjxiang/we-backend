package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"we-backend/pkg/consts"
)

// 测试
func (h *AuthHandler) AuthzMe(c *gin.Context) {

	userId := c.MustGet(consts.SessionKeyUserId).(int64)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "用户编号",
		"data":  userId,
	})
}
