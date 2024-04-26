package routes

import (
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)


func UserRoutes(cfg *config.Config, apiv1 *gin.RouterGroup, userHandler *handler.UserHandler, authHandler *handler.AuthHandler, middleware middleware.Middleware) {

	// 默认设置 gin.DefaultWriter = os.Stdout
	apiv1.Use(gin.Logger())
	// 从任何 panic 恢复，并返回一个 500 错误
	apiv1.Use(gin.Recovery())
	// 跨域
	apiv1.Use(middleware.Cors())
	
	// 健康检查
	apiv1.GET("health", userHandler.HealthCheck)
	
	userGroup := apiv1.Group("/user")
	
	if cfg.EnableDeploySingle() {
		// 认证（cookie and session Part 1）
		store := cookie.NewStore([]byte(cfg.GetSecretKey()))
		// Also set Secure: true if using SSL, you should though
		store.Options(sessions.Options{
				HttpOnly: true, 
				MaxAge:   7 * 86400,  // 7 天
				Path:     "/",
			})
		userGroup.Use(sessions.Sessions("gin-session", store))

		// 认证（cookie and session Part 2）
		userGroup.Use(middleware.AuthenticateUserByCookieAndSession("/api/v1/user/signup", "/api/v1/user/login"))
	} else {
		store, err := redis.NewStore(
			16, 
			cfg.RedisNetworkType, 
			cfg.GetRedisAddr(), 
			cfg.RedisPassword, 
			[]byte(cfg.GetSecretKey()), 
			[]byte(cfg.GetSecretKey()),
		)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to new redis store")
		}

		userGroup.Use(sessions.Sessions("gin-session", store))
		userGroup.Use(middleware.AuthenticateUserByCookieAndSession("/api/v1/user/signup", "/api/v1/user/login"))
	}

	{
		// Todo. 如果了解 gin 注册`中间件和路由`顺序的骚操作，就意识到多此一举
		userGroup.POST("/signup", userHandler.UserSignup)
		userGroup.POST("/login", userHandler.UserLogin)
		userGroup.POST("/edit", userHandler.Edit)
		userGroup.GET("/profile", userHandler.UserProfile)
		userGroup.GET("/logout", userHandler.UserLogout)	
	}
	
}
