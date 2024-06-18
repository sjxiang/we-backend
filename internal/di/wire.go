package di

import (
	"we-backend/internal/api"
	"we-backend/internal/api/handler"
	"we-backend/internal/api/middleware"
	"we-backend/internal/biz"
	"we-backend/internal/conf"
	"we-backend/internal/data"
	"we-backend/internal/service/accesscontrol"
	"we-backend/internal/service/token"
)

func InitializeApi(cfg *conf.Config) (*api.Server, error) {

	// db cache  
	db := data.NewDB(cfg)
	cache := data.NewCache(cfg)

	// repository
	userDB := data.NewUserDatabase(db)
	userCache := data.NewUserCache(cache, cfg)
	userRepo := data.NewUseRepo(userDB, userCache)

	// external service
	tokenService := token.NewTokenService(cfg)
	rateLimitService := accesscontrol.NewRateLimitService(cache, cfg)

	// usecase
	userUsecase := biz.NewUserUsecase(userRepo, tokenService)

	// handler
	handler := handler.NewHandler(userUsecase)

	// middleware
	middleware := middleware.NewMiddleware(tokenService, rateLimitService)

	return api.NewServer(cfg, handler, middleware), nil 
}

