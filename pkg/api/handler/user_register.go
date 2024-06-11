package handler

import (
	"context"
	"encoding/json"

	"we-backend/pkg/errno"
	"we-backend/pkg/types"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"

	"github.com/gin-gonic/gin"
)


func (h *handler) Register(c *gin.Context) {
	var req types.RegisterRequest
	
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrMissingParameter.WithMessage(err.Error()))
		return
	}
	
	if err := validate.Check(req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	minSize, digit, special, letter := utils.ValidatePassword(req.Password)
	if !minSize || !digit || !special || !letter {
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage("密码必须包含数字（123...）、字母（aA...）、特殊字符（@#$...），并且长度不能小于 8 位"))
		return
	}

	rsp, err := h.userUsecase.UserRegister(context.Background(), &req)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c, rsp)
}