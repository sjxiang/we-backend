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
	
	// db
	db := data.NewDB(cfg)
	
	// repository
	userRepo := data.NewUserRepo(db)

	// usecase
	userUsecase := biz.NewUserUsecase(userRepo)

	// handler
	userHandler := handler.NewUserHandler(userUsecase)
	authHandler := handler.NewAuthHandler(userUsecase)

	// middleware
	middleware := middleware.NewMiddleware()

	return api.NewServerHTTP(cfg, userHandler, authHandler, middleware), nil 
}