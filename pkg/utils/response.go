package utils

import (
	"net/http"
	"we-backend/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}


func FeedbackOK(c *gin.Context, data interface{}) {
	resp := Response{
		Data: data,
	}

	c.JSON(http.StatusOK, resp)
}

func FeedbackBadRequest(c *gin.Context, err error) {
	e := errno.ConvertErr(err)

	c.JSON(http.StatusBadRequest, Response{
		Code:    e.ErrCode,
		Message: e.ErrMsg,
	})
}

func FeedbackInternalServerError(c *gin.Context, err error) {
	e := errno.ConvertErr(err)

	c.JSON(http.StatusInternalServerError, Response{
		Code:    e.ErrCode,
		Message: e.ErrMsg,
	})
}
