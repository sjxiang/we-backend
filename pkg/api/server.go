package api

import (
	"net/http"

	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/api/routes"
	"we-backend/pkg/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)



type ServerHTTP struct {
	Engine *gin.Engine
}

func NewServerHTTP(cfg *config.Config, userHandler *handler.UserHandler, authHandler *handler.AuthHandler, middleware middleware.Middleware) *ServerHTTP {

	engine := gin.New()

	// no handler
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404 page not found",
		})
	})
	engine.NoMethod(func(c *gin.Context) {
		c.Data(http.StatusMethodNotAllowed, "text/plain", []byte("Method Not Allowed"))
		c.Abort()
	})

	// set up routes
	routes.UserRoutes(engine.Group("/api/v1"), userHandler, authHandler, middleware)
	
	if cfg.EnableAuthzBySession() {
		// set up session (Part 1)
		store := cookie.NewStore([]byte("8xEMrWkBARcDDYQ"))
		// Also set Secure: true if using SSL, you should though
		store.Options(sessions.Options{
				HttpOnly: true, 
				MaxAge:   7 * 86400,  // 7 天
				Path:     "/",
			})
		engine.Use(sessions.Sessions("gin-session", store))

		// set up session (Part 3)
		engine.Use(middleware.AuthenticateUserByCookieAndSession())
	}

	if cfg.EnableAuthzByJWT() {
		panic("Implement me!")
	}

	return &ServerHTTP{Engine: engine}
}


func (s *ServerHTTP) Start() error {
	return s.Engine.Run(":8000")
}
