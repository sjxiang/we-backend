package handler

import (
	"encoding/json"

	"we-backend/pkg/errno"
	"we-backend/pkg/types"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"


	"github.com/gin-gonic/gin"
)

// 编辑用户信息
func (h *handler) EditUser(c *gin.Context) {
	var req types.EditInfoRequest

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