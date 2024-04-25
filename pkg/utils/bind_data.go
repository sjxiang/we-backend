package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type invalidArgument struct {
	Field string `json:"field"`  // 字段
	Value string `json:"value"`  // 值
	Tag   string `json:"tag"`    // binding tag
	Param string `json:"param"`  
}

func BindData(c *gin.Context, req interface{}) bool {

	if c.ContentType() != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": true,
			"msg":   msg,
		})
		return false
	}

	if err := c.ShouldBind(req); err != nil {
		
		// errs 是一个错误 slice，保存参数校验失败的字段和错误信息（不满足 binding tag 条件）
		errs, ok := err.(validator.ValidationErrors)
		
		if ok {
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),            
					err.Value().(string),   
					err.Tag(),              
					err.Param(),            
				})
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"msg":   fmt.Sprintf("Invalid request parameters, See invalidArgs: %+v", invalidArgs),
			})
			return false
		}

		// 其他错误（1. 参数绑定失败，类型不匹配 json.Unmarshal、2. 服务器内部错误 internal server error）
		fallBack := fmt.Errorf("other error happened %w", err).Error()

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   fallBack,
		}) 
		return false
	}

	return true
}
