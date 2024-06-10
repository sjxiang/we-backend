package api

import (
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/api/routes"
	"we-backend/pkg/config"

	"github.com/gin-gonic/gin"
)



type ServerHTTP struct {
	Engine *gin.Engine
}

func NewServerHTTP(cfg *config.Config, userHandler *handler.UserHandler, middleware middleware.Middleware) *ServerHTTP {

	engine := gin.New()

	// 404
	engine.NoRoute(func(c *gin.Context) {
		c.Data(404, "text/plain", []byte("404 page not found"))
		c.Abort()
	})
	engine.NoMethod(func(c *gin.Context) {
		c.Data(405, "text/plain", []byte("Method Not Allowed"))
		c.Abort()
	})


	// set up routes
	routes.UserRoutes(cfg, engine.Group("/api/v1"), userHandler, middleware)


	return &ServerHTTP{Engine: engine}
}


func (s *ServerHTTP) Start() error {
	return s.Engine.Run(":5678")
}
