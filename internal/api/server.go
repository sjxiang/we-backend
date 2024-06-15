package api

import (
	"github.com/gin-gonic/gin"

	"we-backend/internal/api/handler"
	"we-backend/internal/api/middleware"
	"we-backend/internal/api/routes"
	"we-backend/internal/conf"
)


type Server struct {
	engine *gin.Engine
	cfg    *conf.Config
}

func NewServer(cfg *conf.Config, h handler.Handler, m middleware.Middleware) *Server {

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
	
	// 默认设置 gin.DefaultWriter = os.Stdout、发生 painc 返回一个 500、跨域、滑动窗口限流
	s.engine.Use(gin.Logger(), gin.Recovery(), m.EnableCORS())

	routes.UserRoutes(s.engine.Group("/api/v1/user"), h, m)
}