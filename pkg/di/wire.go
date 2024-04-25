package di

import (
	"we-backend/pkg/api"
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/biz"
	"we-backend/pkg/config"
	"we-backend/pkg/data"
)

func InitializeApi(cfg *config.Config) (*api.ServerHTTP, error) {
	
	// external service
	
	// repository
	db := data.NewDB(cfg)
	ur := data.NewUserRepo(db)

	// usecase
	uc := biz.NewUserUsecase(ur)

	// handler
	userHandler := handler.NewUserHandler(uc)
	authHandler := handler.NewAuthHandler(uc)

	// middleware
	middleware := middleware.NewMiddleware()

	return api.NewServerHTTP(cfg, userHandler, authHandler, middleware), nil 
}