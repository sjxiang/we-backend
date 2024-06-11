package routes

import (
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	
	"github.com/gin-gonic/gin"
)


func UserRoutes(apiv1 *gin.RouterGroup, h handler.Handler, m middleware.Middleware) {
	
	userGroup := apiv1.Group("/user")
	{
		userGroup.GET("/health", h.HealthCheck)
		userGroup.POST("/register", h.Register)
		userGroup.POST("/login", h.Login)
		
		// 了解 gin 注册`中间件和路由`顺序的骚操作
		userGroup.Use(m.Authenticate())
		
		userGroup.GET("/me", h.Me)
		userGroup.POST("/edit", h.EditUser)
	}

}
