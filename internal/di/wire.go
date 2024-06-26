package di

import (
	"we-backend/internal/api"
	"we-backend/internal/api/handler"
	"we-backend/internal/api/middleware"
	"we-backend/internal/biz"
	"we-backend/internal/conf"
	"we-backend/internal/data"
	"we-backend/internal/service/accesscontrol"
	"we-backend/internal/service/mail"
	"we-backend/internal/service/token"
)

func InitializeApi(cfg *conf.Config) (*api.Server, error) {

	// db cache  
	db := data.NewDB(cfg)
	cache := data.NewCache(cfg)

	// repository
	userDB := data.NewUserDatabase(db)
	userCache := data.NewUserCache(cache, cfg)
	userRepo := data.NewUserRepo(userDB, userCache)
	otpCache := data.NewOtpCache(cache)
	otpRepo := data.NewOtpRepo(otpCache)

	// external service
	tokenService := token.NewTokenService(cfg)
	accessControlService := accesscontrol.NewAccessControlService(cache, cfg)
	emailService := mail.NewQQMailSender(cfg)

	// usecase
	userUsecase := biz.NewUserUsecase(userRepo, tokenService)
	otpUsecase := biz.NewOtpUsecase(otpRepo, emailService)

	// handler
	handler := handler.NewHandler(userUsecase, otpUsecase)

	// middleware
	middleware := middleware.NewMiddleware(tokenService, accessControlService)

	return api.NewServer(cfg, handler, middleware), nil 
}

