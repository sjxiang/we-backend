package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查看用户详情 detail
func (h *UserHandler) UserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "测试成功",
	})
}
