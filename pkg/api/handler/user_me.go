package handler

import (
	"we-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 查看用户详情
func (h *UserHandler) Me(c *gin.Context) {

	// userID := c.MustGet(consts.SessionKeyUserId).(int64)

	
	utils.FeedbackOK(c, "")
}
