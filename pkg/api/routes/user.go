package routes

import (
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	
	"github.com/gin-gonic/gin"
)


func UserRoutes(apiv1 *gin.RouterGroup, userHandler *handler.UserHandler, authHandler *handler.AuthHandler, middleware middleware.Middleware) {

	// 默认设置 gin.DefaultWriter = os.Stdout
	apiv1.Use(gin.Logger())
	// 从任何 panic 恢复，并返回一个 500 错误
	apiv1.Use(gin.Recovery())
	// 跨域
	apiv1.Use(middleware.HandleCors())
	
	userGroup := apiv1.Group("/user")
	
	{
		userGroup.POST("/signup", userHandler.UserSignup)
		userGroup.POST("/login", userHandler.UserLogin)
	}

	userGroup.Use()

	{	
		userGroup.POST("/edit", userHandler.Edit)
		userGroup.GET("/profile", userHandler.UserProfile)
		userGroup.GET("/logout", userHandler.Logout)
	}
	

}
