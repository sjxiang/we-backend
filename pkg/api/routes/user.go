package routes

import (
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/config"

	"github.com/gin-gonic/gin"
)


func UserRoutes(cfg *config.Config, apiv1 *gin.RouterGroup, userHandler *handler.UserHandler, authHandler *handler.AuthHandler, middleware middleware.Middleware) {

	// 默认设置 gin.DefaultWriter = os.Stdout
	apiv1.Use(gin.Logger())
	// 从任何 panic 恢复，并返回一个 500 错误
	apiv1.Use(gin.Recovery())
	// 跨域
	apiv1.Use(middleware.Cors())
	


	// *** 健康检查 ***
	apiv1.GET("health", authHandler.HealthCheck)
	

	// *** 认证 ***
	authGroup := apiv1.Group("/authz")


	// cookie
	if cfg.EnableAuthzCookie() {
		ignorePaths := make([]string, 0)
		
		ignorePaths = append(ignorePaths, "/api/v1/authz/signup")
		ignorePaths = append(ignorePaths, "/api/v1/authz/login")
		authGroup.Use(middleware.AuthenticateUserByCookie(ignorePaths...))

		// Todo. 如果了解 gin 注册`中间件和路由`顺序的骚操作，就意识到多此一举
		authGroup.POST("/signup", authHandler.AuthzSignup)
		authGroup.POST("/login", authHandler.AuthzLoginByCookie)
		authGroup.GET("/me", authHandler.AuthzMe)
	}

	
	// *** 用户 ***
	userGroup := apiv1.Group("/user")
	{
		// Todo. 如果了解 gin 注册`中间件和路由`顺序的骚操作，就意识到多此一举
		userGroup.POST("/edit", userHandler.EditUser)
		userGroup.GET("/profile", userHandler.UserProfile)
	}
	
}
