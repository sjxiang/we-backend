package di

import (
	"we-backend/internal/api"
	"we-backend/internal/api/handler"
	"we-backend/internal/api/middleware"
	"we-backend/internal/biz"
	"we-backend/internal/conf"
	"we-backend/internal/data"
	"we-backend/internal/service/access"
	"we-backend/internal/service/token"
)

func InitializeApi(cfg *conf.Config) (*api.Server, error) {

	// db
	db := data.NewDB(cfg)
	cache := data.NewCache(cfg)
	
	// external service
	tokenService := token.NewTokenService(cfg)
	accessControlService := access.NewAccessControlService(cache, cfg)

	// repository
	userRepo := data.NewUserRepo(db)
	userCache := data.NewUserCache(cache, cfg)

	// usecase
	userUsecase := biz.NewUserUsecase(userRepo, userCache, tokenService)

	// handler
	handler := handler.NewHandler(userUsecase)

	// middleware
	middleware := middleware.NewMiddleware(tokenService, accessControlService)

	return api.NewServer(cfg, handler, middleware), nil 
}

