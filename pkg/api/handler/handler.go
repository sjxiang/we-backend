package handler

import "we-backend/pkg/biz"


type (	
	// 认证相关
	AuthHandler struct {
		uc *biz.UserUsecase

	}

	// 用户相关
	UserHandler struct {
		usecase *biz.UserUsecase
	}
)

func NewAuthHandler(uc *biz.UserUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

func NewUserHandler(uc *biz.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}
