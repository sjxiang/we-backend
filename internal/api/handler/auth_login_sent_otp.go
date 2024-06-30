package handler

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/we"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"
)

func (h *handler) LoginBySentOtp(c *gin.Context) {
	
	var req types.SentOtpRequest
	
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrMissingParameter.WithMessage(err.Error()))
		return
	}

	if err := validate.Check(req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	input := types.SentOtpInput{
		Biz:         "login",
		PhoneNumber: req.PhoneNumber,
	}
	if err := h.AuthUsecase.LoginBySentOtp(context.Background(), input); err != nil {
		utils.FeedbackBadRequest(c, err)
		return 
	}

	utils.FeedbackOK(c, "验证码已发送至你的手机", nil)
}