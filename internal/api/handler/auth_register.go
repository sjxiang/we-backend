package handler

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/faker"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"
	"we-backend/pkg/we"
)


func (h *handler) Register(c *gin.Context) {
	var req types.RegisterRequest
	
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrMissingParameter.WithMessage(err.Error()))
		return
	}
	
	if err := validate.Check(req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	input := types.RegisterInput{
		Username:        faker.Username(),
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.PasswordConfirm,
	}
	resp, err := h.AuthUsecase.Register(context.Background(), input)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c, "用户注册成功", resp)
}