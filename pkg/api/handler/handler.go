package handler

import "we-backend/pkg/biz"


type UserHandler struct {
	usecase *biz.UserUsecase
}

func NewUserHandler(userUsecase *biz.UserUsecase) *UserHandler {
	return &UserHandler{usecase: userUsecase}
}

