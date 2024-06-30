package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/we"
	"we-backend/pkg/utils"
)

func (h *handler) Login(c *gin.Context) {

	var req types.LoginRequest
	
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrMissingParameter.WithMessage(err.Error()))
		return
	}

	if violations := req.Validate(); len(violations) > 0 {
		utils.FeedbackBadRequest(c, we.ErrInvalidParameter.WithMessage(invalidArgumentError(violations).Error()))
		return
	}

	input := types.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}
	resp, err := h.AuthUsecase.Login(context.Background(), input)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c, "用户登录成功", resp)
}	

func invalidArgumentError(violations []string) error {
	details := strings.Join(violations, "、")
	return fmt.Errorf("field violations: %s", details)
}
