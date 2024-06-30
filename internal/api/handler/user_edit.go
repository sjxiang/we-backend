package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/we"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"
)

// 编辑用户信息
func (h *handler) Edit(c *gin.Context) {
	userID, err := utils.GetUserIDFromAuth(c)
	if err != nil {
		utils.FeedbackBadRequest(c, we.ErrNotLogin.WithMessage(err.Error()))
		return
	}

	var req types.EditRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrMissingParameter.WithMessage(err.Error()))
		return
	}

	if err := validate.Check(req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	birthday, err := time.Parse(time.DateOnly, req.Birthday)
	if err != nil {
		utils.FeedbackBadRequest(c, we.ErrInvalidParameter.WithMessage(err.Error()))  // 生日格式不对 YYYY-MM-DD
		return  
	}

	input := types.EditInput{
		UserID:   userID,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Intro:    req.Intro,
		Birthday: birthday.Unix(), // 将时间转换为时间戳
	}

	if err := h.UserUsecase.Edit(context.Background(), input); err != nil {
		utils.FeedbackBadRequest(c, err)
		return 
	}

	utils.FeedbackOK(c, "编辑用户信息成功", nil)
}