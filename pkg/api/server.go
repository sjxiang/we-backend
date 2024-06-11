package api

import (
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/api/routes"
	"we-backend/pkg/config"

	"github.com/gin-gonic/gin"
)


type Server struct {
	engine *gin.Engine
	cfg    *config.Config
}

func NewServer(cfg *config.Config, h handler.Handler, m middleware.Middleware) *Server {

	server := Server{
		engine: gin.New(),
		cfg:    cfg,
	}

	server.routes(h, m)

	return &server
}


func (s *Server) Start() error {
	return s.engine.Run(":5678")
}


func (s *Server) routes(h handler.Handler, m middleware.Middleware) {
	// 404
	s.engine.NoRoute(func(c *gin.Context) {
		c.Data(404, "text/plain", []byte("404 page not found"))
		c.Abort()
	})
	s.engine.NoMethod(func(c *gin.Context) {
		c.Data(405, "text/plain", []byte("Method Not Allowed"))
		c.Abort()
	})
	
	// 默认设置 gin.DefaultWriter = os.Stdout、发生 painc 返回一个 500、跨域
	s.engine.Use(gin.Logger(), gin.Recovery(), m.EnableCORS())

	apiv1 := s.engine.Group("/api/v1")

		
	routes.UserRoutes(apiv1, h, m)
}