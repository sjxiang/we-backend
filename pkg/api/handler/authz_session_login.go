package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"we-backend/pkg/consts"
	"we-backend/pkg/types"
	"we-backend/pkg/x"
)


func (h *AuthHandler) AuthzLoginBySession(c *gin.Context) {
	// fetch payload
	req := types.NewLoginRequest()

	if err := json.NewDecoder(c.Request.Body).Decode(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   fmt.Sprintf("parse request body error\n %+v", err),  
		})
		return
	}

	// validate payload required fields
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   fmt.Sprintf("validate request body error\n %+v", err),
		})
		return
	}

	// handle biz 
	user, err := h.usecase.UserLogin(
		context.TODO(), 
		req.ExportEmailInString(), 
		req.ExportPasswordInString(),
	)
	if err != nil {
		if errors.Is(err, x.ErrInvalidCredentials) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"msg":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	// 设置 session
	var (
		uid int64     = user.ID
		now time.Time = time.Now()
	)
	
	session := sessions.Default(c)
	session.Set(consts.SessionKeyUserId, uid)
	session.Set(consts.SessionKeyLastTime, now)
	session.Save()

	// feedback
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "登录成功",
	})
} 