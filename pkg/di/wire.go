package di

import (
	"we-backend/pkg/api"
	"we-backend/pkg/api/handler"
	"we-backend/pkg/api/middleware"
	"we-backend/pkg/biz"
	"we-backend/pkg/config"
	"we-backend/pkg/data"
	"we-backend/pkg/service/token"
)

func InitializeApi(cfg *config.Config) (*api.Server, error) {
	
	// external service
	tokenService := token.NewTokenService(cfg)

	// db
	db := data.NewDB(cfg)
	
	// repository
	userRepo := data.NewUserRepo(db)

	// usecase
	userUsecase := biz.NewUserUsecase(userRepo, tokenService)

	// handler
	handler := handler.NewHandler(userUsecase)

	// middleware
	middleware := middleware.NewMiddleware(tokenService)

	return api.NewServer(cfg, handler, middleware), nil 
}