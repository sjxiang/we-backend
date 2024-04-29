package routes

import (
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)


func UserRoutes(cfg *config.Config, apiv1 *gin.RouterGroup, userHandler *handler.UserHandler, authHandler *handler.AuthHandler, middleware middleware.Middleware) {

	// 默认设置 gin.DefaultWriter = os.Stdout
	apiv1.Use(gin.Logger())
	// 从任何 panic 恢复，并返回一个 500 错误
	apiv1.Use(gin.Recovery())
	// 跨域
	// apiv1.Use(middleware.Cors())
	
	// 健康检查
	apiv1.GET("health", authHandler.HealthCheck)
	

	// 认证
	authGroup := apiv1.Group("/authz")

	// cookie
	if cfg.EnableAuthzCookie() {
		authGroup.Use(middleware.AuthenticateUserByCookie(
			"/api/v1/authz/signup", 
			"/api/v1/authz/login", 
			))
		authGroup.POST("/login", authHandler.AuthzLoginByCookie)
	}

	// single session 
	if cfg.EnableAuthzSingleSession() {

		store := cookie.NewStore([]byte(cfg.GetSecretKey()))
		// Also set Secure: true if using SSL, you should though
		store.Options(sessions.Options{
			HttpOnly: true, 
			MaxAge:   7 * 86400,  // 7 天
			Path:     "/",
		})
		authGroup.Use(middleware.AuthenticateUserByCookieAndSession(
			"/api/v1/authz/signup", 
			"/api/v1/authz/login", 
			))
	
		authGroup.POST("/login", authHandler.AuthzLoginBySession)
	}

	// multi session
	if cfg.EnableAuthzMultiSession() {
		panic("implement me")
	}

	// jwt
	if cfg.EnableAuthzJWT() {
		panic("implement me")
	}
	
	{
		authGroup.POST("/signup", authHandler.AuthzSignup)
	}

	
	userGroup := apiv1.Group("/user")
	{
		// Todo. 如果了解 gin 注册`中间件和路由`顺序的骚操作，就意识到多此一举
		userGroup.POST("/edit", userHandler.Edit)
		userGroup.GET("/profile", userHandler.UserProfile)
	}
	
}
