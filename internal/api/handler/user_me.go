package handler

import (
	"context"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/utils"
	"we-backend/pkg/we"
)

// 查看用户详情
func (h *handler) Me(c *gin.Context) {

	userID, err := utils.GetUserIDFromAuth(c)
	if err != nil {
		utils.FeedbackBadRequest(c, we.ErrNotLogin.WithMessage(err.Error()))
		return
	}

	input := types.MeInput{
		UserID: userID,
	}
	resp, err := h.UserUsecase.Me(context.Background(), input)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c, "用户详情", resp)
}
