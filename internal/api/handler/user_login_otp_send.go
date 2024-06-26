package handler

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/errno"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"
)

func (h *handler) SentOtp(c *gin.Context) {
	var req types.SentOtpRequest
	
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrMissingParameter.WithMessage(err.Error()))
		return
	}

	if err := validate.Check(req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	const biz = "login"
	if err := h.otpUsecase.SendOtp(context.TODO(), biz, req.PhoneNumber); err != nil {
		utils.FeedbackBadRequest(c, err)
		return 
	}

	utils.FeedbackOK(c, "发送成功")
}