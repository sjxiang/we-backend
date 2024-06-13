package di

import (
	"we-backend/internal/api"
	"we-backend/internal/api/handler"
	"we-backend/internal/api/middleware"
	"we-backend/internal/biz"
	"we-backend/internal/conf"
	"we-backend/internal/data"
	"we-backend/internal/service/limiter"
	"we-backend/internal/service/token"
)

func InitializeApi(cfg *conf.Config) (*api.Server, error) {

	// db
	db := data.NewDB(cfg)
	cache := data.NewCache(cfg)
	
	// external service
	tokenService := token.NewTokenService(cfg)
	rateLimitService := limiter.NewRateLimitService(cache, cfg)

	// repository
	userRepo := data.NewUserRepo(db)

	// usecase
	userUsecase := biz.NewUserUsecase(userRepo, tokenService)

	// handler
	handler := handler.NewHandler(userUsecase)

	// middleware
	middleware := middleware.NewMiddleware(tokenService, rateLimitService)

	return api.NewServer(cfg, handler, middleware), nil 
}