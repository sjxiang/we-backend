package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"we-backend/pkg/types"
	"we-backend/pkg/x"
)

const DEFAULT_COOKIE_NAME      = "_cookie"
const DEFAULT_COOKIE_MAX_AGE   = 3600         // cookie 的有效期，缺省时，cookie 仅在浏览器关闭之前有效
const DEFAULT_COOKIE_PATH      = "/"          // 限制指定 cookie 的发送范围的文件目录，默认为当前
const DEFAULT_COOKIE_DOMAIN    = "localhost"  // 限制 cookie 生效的域名，默认为创建 cookie 的服务域
const DEFAULT_COOKIE_SECURE    = false        // 仅在 HTTPS 安全连接时，才可以发送 cookie
const DEFAULT_COOKIE_HTTP_ONLY = true         // Javascript 脚本无法获得 cookie

// 登录
func (h *AuthHandler) AuthzLoginByCookie(c *gin.Context) {

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

	// 设置和页面关联的 cookie，服务器通过这个头部把 cookie 传给客户端
	c.SetCookie(
		DEFAULT_COOKIE_NAME,
		fmt.Sprintf("%d", user.ID),
		DEFAULT_COOKIE_MAX_AGE, 
		DEFAULT_COOKIE_PATH, 
		DEFAULT_COOKIE_DOMAIN, 
		DEFAULT_COOKIE_SECURE, 
		DEFAULT_COOKIE_HTTP_ONLY,
	)

	// feedback
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "登录成功",
	})
}
