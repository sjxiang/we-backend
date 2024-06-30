package handler

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"we-backend/internal/types"
	"we-backend/pkg/utils"
	"we-backend/pkg/validate"
	"we-backend/pkg/we"
)

func (h *handler) LoginByVerifyOtp(c *gin.Context) {

	var req types.VerifyOtpRequest
	
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrMissingParameter.WithMessage(err.Error()))
		return
	}

	if err := validate.Check(req); err != nil {
		utils.FeedbackBadRequest(c, we.ErrInvalidParameter.WithMessage(err.Error()))
		return
	}

	input := types.VerifyOtpInput{
		Biz:         "login",
		PhoneNumber: req.PhoneNumber,
		InputCode:   req.InputCode,
	}
	resp, err := h.AuthUsecase.LoginByVerifyOtp(context.TODO(), input)
	if err != nil {
		utils.FeedbackBadRequest(c, err)
		return 
	}

	utils.FeedbackOK(c, "验证码校验通过", resp)
}

/*


更多登录方式

1， 切换到手机号验证

输入手机号验证码

	请输入发送至 +86ooo****xxxx 的 6 位验证码，有效期 10 分钟
		xxx - xxx
		
	60 秒后可重新获取验证码


2. 切换到密码验证
	
忘记密码？点此重置


 */