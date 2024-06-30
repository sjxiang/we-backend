package handler

import (
	"we-backend/internal/biz"

	"github.com/gin-gonic/gin"
)


type handler struct {
	UserUsecase  *biz.UserUsecase
	AuthUsecase  *biz.AuthUsecase
}

func NewHandler(ur *biz.UserUsecase, au *biz.AuthUsecase) Handler {
	return &handler{
		UserUsecase: ur,
		AuthUsecase: au,
	}
}



type Handler interface {
	HealthCheck(c *gin.Context)

	Register(c *gin.Context) 
	Login(c *gin.Context)
	SentOtp(c *gin.Context)
	VerifyOtp(c *gin.Context)

	Me(c *gin.Context)
	Edit(c *gin.Context)
	Admin(c *gin.Context)
	ResetPassword(c *gin.Context) 
}