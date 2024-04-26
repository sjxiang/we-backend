package api

import (
	"net/http"

	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/api/routes"
	"we-backend/pkg/config"

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
	routes.UserRoutes(cfg, engine.Group("/api/v1"), userHandler, authHandler, middleware)


	return &ServerHTTP{Engine: engine}
}


func (s *ServerHTTP) Start() error {
	return s.Engine.Run(":8000")
}
