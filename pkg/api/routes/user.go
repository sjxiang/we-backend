package routes

import (
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/config"

	"github.com/gin-gonic/gin"
)


func UserRoutes(cfg *config.Config, apiv1 *gin.RouterGroup, userHandler *handler.UserHandler, middleware middleware.Middleware) {

	// 默认设置 gin.DefaultWriter = os.Stdout、发生 painc 返回一个 500
	apiv1.Use(gin.Logger(), gin.Logger())
	// 跨域
	apiv1.Use(middleware.EnableCORS())
	
	userGroup := apiv1.Group("/user")
	{
		userGroup.GET("/health", userHandler.HealthCheck)
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
		
		// 了解 gin 注册`中间件和路由`顺序的骚操作
		
		userGroup.GET("/me", userHandler.Me)
		userGroup.POST("/edit", userHandler.EditUser)
	}

}
