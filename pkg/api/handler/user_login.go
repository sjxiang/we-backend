package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"we-backend/pkg/errno"
	"we-backend/pkg/types"
	"we-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) Login(c *gin.Context) {

	var req types.LoginRequest
	
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrMissingParameter.WithMessage(err.Error()))
		return
	}

	if violations := req.Validate(); len(violations) > 0 {
		msg := utils.Mix(violations)
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage(msg))
		return
	}

	fmt.Println(req)
	
	rsp, err := h.userUsecase.UserLogin(context.TODO(), &req)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return
	}

	utils.FeedbackOK(c, rsp.ExportForFeedback())
}	