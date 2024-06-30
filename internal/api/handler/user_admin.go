package handler

import (
	"context"
	"we-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) Admin(c *gin.Context) {
	resp, err := h.UserUsecase.All(context.Background())
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c, "用户列表", resp)
}

