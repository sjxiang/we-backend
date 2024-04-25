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

	"we-backend/pkg/api/handler/request"
	"we-backend/pkg/consts"
	"we-backend/pkg/x"
)

// 用户登录
func (h *UserHandler) UserLogin(c *gin.Context) {

	// fetch payload
	req := request.NewLoginRequest()

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

	// set up session (Part 2)
	s := sessions.Default(c)
	s.Clear()
	s.Set(consts.SessionKeyUserId, user.ID)
	s.Set(consts.SessionKeyLastTime, time.Now())
	s.Save()
	
	// feedback
	c.JSON(http.StatusOK, nil)
}

