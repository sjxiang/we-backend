package handler

import (
	"context"

	"we-backend/pkg/errno"
	"we-backend/pkg/types"
	"we-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)


func (h *UserHandler) Register(c *gin.Context) {
	var req types.RegisterRequest
	
	if err := c.ShouldBind(&req); err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}	
	
	minSize, digit, special, letter := utils.ValidatePassword(req.Password)
	if !minSize || !digit || !special || !letter {
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage("密码必须包含数字（123...）、字母（aA...）、特殊字符（@#$...），并且长度不能小于 8 位"))
		return
	}

	resp, err := h.usecase.UserRegister(context.Background(), &req)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c, resp)
}