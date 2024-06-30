package routes

import (
	"github.com/gin-gonic/gin"

	"we-backend/internal/api/handler"
	"we-backend/internal/api/middleware"
)

func UserRoutes(group *gin.RouterGroup, h handler.Handler, m middleware.Middleware) {
	
		group.GET("/health", h.HealthCheck)
		
		group.POST("/register", h.Register)
		group.POST("/login", h.Login)
		group.POST("/login_sms/otp/send", h.LoginBySentOtp)
		group.POST("/login_sms/otp/verify", h.LoginByVerifyOtp)
		
		
		// 了解 gin 注册`中间件和路由`顺序的骚操作
		group.Use(m.Authenticate())
		
		group.GET("/me", h.Me)
		group.POST("/edit", h.Edit)
		group.GET("/admin", h.Admin)
}
