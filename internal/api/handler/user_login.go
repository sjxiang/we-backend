package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/errno"
	"we-backend/pkg/utils"
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