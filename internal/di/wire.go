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
	captchaRepo := data.NewCaptchaRepo(cache)

	// external service
	tokenService := token.NewTokenService(cfg)
	accessControlService := accesscontrol.NewAccessControlService(cache, cfg)
	emailService := mail.NewLocalMailSender(cfg)

	// usecase
	authUsecase := biz.NewAuthUsecase(userRepo, captchaRepo, tokenService, emailService)
	userUsecase := biz.NewUserUsecase(userRepo)
	
	// handler
	handler := handler.NewHandler(userUsecase, authUsecase)

	// middleware
	middleware := middleware.NewMiddleware(tokenService, accessControlService)

	return api.NewServer(cfg, handler, middleware), nil 
}

