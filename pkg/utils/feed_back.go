package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"we-backend/pkg/we"
)

type Response struct {
	Code    uint32      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}


func FeedbackOK(c *gin.Context, msg string, data any) {
	resp := Response{
		Message: msg,
		Data:    data,
	}

	c.JSON(http.StatusOK, resp)
}

func FeedbackBadRequest(c *gin.Context, err error) {
	e := we.ConvertErr(err)

	c.JSON(http.StatusBadRequest, Response{
		Code:    e.ErrCode,
		Message: e.ErrMsg,
	})
}

func FeedbackInternalServerError(c *gin.Context, err error) {
	e := we.ConvertErr(err)

	c.JSON(http.StatusInternalServerError, Response{
		Code:    e.ErrCode,
		Message: e.ErrMsg,
	})
}


func FeedbackAuthorizedFailedError(c *gin.Context, err error) {
	e := we.ConvertErr(err)

	c.AbortWithStatusJSON(http.StatusUnauthorized, Response{
		Code:    e.ErrCode,
		Message: e.ErrMsg,
	})
}
