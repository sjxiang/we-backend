package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/errno"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"
)

// 编辑用户信息
func (h *handler) Edit(c *gin.Context) {
	var req types.EditRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrMissingParameter.WithMessage(err.Error()))
		return
	}
	
	if err := validate.Check(req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	utils.FeedbackOK(c, nil)
}