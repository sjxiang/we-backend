package handler

import (
	"we-backend/pkg/biz"

	"github.com/gin-gonic/gin"
)


type handler struct {
	userUsecase     *biz.UserUsecase
}

func NewHandler(userUsecase *biz.UserUsecase) Handler {
	return &handler{
		userUsecase: userUsecase, 
	}
}


type Handler interface {
	HealthCheck(c *gin.Context)
	Register(c *gin.Context) 
	Login(c *gin.Context)
	Me(c *gin.Context)
	EditUser(c *gin.Context)
}