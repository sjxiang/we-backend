package handler

import (
	"context"
	"we-backend/pkg/errno"
	"we-backend/pkg/types"
	"we-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 查看用户详情
func (h *handler) Me(c *gin.Context) {

	userID, err := utils.GetUserIDFromAuth(c)
	if err != nil {
		utils.FeedbackBadRequest(c, errno.ErrMissingParameter.WithMessage("请重新登录"))
		return
	}

	req := types.ProfileRequest{
		UserID: userID,
	}

	rsp, err := h.userUsecase.UserProfile(context.TODO(), &req)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c,  rsp.ExportForFeedback())
}
