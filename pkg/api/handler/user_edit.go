package handler

import (
	"we-backend/pkg/errno"
	"we-backend/pkg/types"
	"we-backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 编辑用户信息
func (h *UserHandler) EditUser(c *gin.Context) {
	var req types.EditInfoRequest

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		utils.FeedbackBadRequest(c,errno.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	utils.FeedbackOK(c, nil)
}