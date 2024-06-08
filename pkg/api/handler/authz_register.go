package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"we-backend/pkg/types"
	"we-backend/pkg/utils"
	"we-backend/pkg/x"
)

// 注册 register
func (h *AuthHandler) AuthzSignup(c *gin.Context) {
	
	var req types.SignupRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		return // InvalidJSON
	}
	
	if ok := utils.BindData(c, &req); !ok {
		return
	}
	
	// validate
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "两次密码输入不一致",
		})
		return
	}

	minSize, digit, special, letter := utils.ValidatePasswordMiddle(req.Password)
	if !minSize || !digit || !special || !letter {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "密码必须包含数字（123...）、字母（aA...）、特殊字符（@#$...），并且长度不能小于 8 位",
		})
		return
	}

	// fetch data
	if err := h.usecase.UserSignup(context.TODO(), req.Email, req.Password); err != nil {
		if errors.Is(err, x.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{
				"error": true,
				"msg":   "邮箱已注册",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   "系统异常",
		})
		return
	}

	// feedback
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "注册成功",
	})
}
