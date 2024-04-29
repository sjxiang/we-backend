package handler

import (
	"net/http"
	"we-backend/pkg/consts"

	"github.com/gin-gonic/gin"
)

// 查看用户详情 detail
func (h *UserHandler) UserProfile(c *gin.Context) {

	userId := c.MustGet(consts.SessionKeyUserId).(int64)

	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "测试成功",
		"data":  userId,
	})
}
