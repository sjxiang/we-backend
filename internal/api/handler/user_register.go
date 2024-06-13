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
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage("这个密码太弱了，不少于8个字符，必须包含大写和小写字母、数字以及特殊符号"))
		return
	}

	rsp, err := h.userUsecase.UserRegister(context.Background(), &req)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c, rsp)
}