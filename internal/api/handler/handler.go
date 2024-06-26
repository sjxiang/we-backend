package handler

import (
	"we-backend/internal/biz"

	"github.com/gin-gonic/gin"
)


type handler struct {
	userUsecase     *biz.UserUsecase
	otpUsecase      *biz.OtpUsecase
}

func NewHandler(userUsecase *biz.UserUsecase, otpUsecase *biz.OtpUsecase) Handler {
	return &handler{
		userUsecase: userUsecase, 
		otpUsecase:  otpUsecase,
	}
}


type Handler interface {
	HealthCheck(c *gin.Context)
	Register(c *gin.Context) 
	Login(c *gin.Context)
	Me(c *gin.Context)
	Edit(c *gin.Context)
	SentOtp(c *gin.Context)
	VerifyOtp(c *gin.Context)
}