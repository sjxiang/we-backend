package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/errno"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"
)

// 编辑用户信息
func (h *handler) Edit(c *gin.Context) {
	userID, err := utils.GetUserIDFromAuth(c)
	if err != nil {
		utils.FeedbackBadRequest(c, errno.ErrMissingParameter.WithMessage("请重新登录"))
		return
	}


	var req types.EditRequest

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrMissingParameter.WithMessage(err.Error()))
		return
	}

	if err := validate.Check(req); err != nil {
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	birthday, err := time.Parse(time.DateOnly, req.Birthday)
	if err != nil {
		utils.FeedbackBadRequest(c, errno.ErrInvalidParameter.WithMessage(err.Error()))  // 生日格式不对 YYYY-MM-DD
		return  
	}

	arg := &types.EditParam{
		UserID:   userID,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Intro:    req.Intro,
		Birthday: birthday.Unix(), // 将时间转换为时间戳
	}

	if err := h.userUsecase.UserEditInfo(context.TODO(), arg); err != nil {
		utils.FeedbackBadRequest(c, err)
		return 
	}

	utils.FeedbackOK(c, "编辑成功")
}