package biz

import (
	"context"
	"errors"

	"we-backend/pkg/errno"
	"we-backend/pkg/types"
	"we-backend/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo UserRepo
}

func NewUserUsecase(userRepo UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uc *UserUsecase) UserRegister(ctx context.Context, req *types.RegisterRequest) (*types.RegisterResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	newUser := types.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	id, err := uc.userRepo.Insert(ctx, newUser)
	if err != nil {
		return nil, err 
	}

	resp := &types.RegisterResponse{
		UserID: id,
	}

	return resp, nil 
}


func (uc *UserUsecase) UserLogin(ctx context.Context, req *types.LoginRequest) (*types.LoginResponse, error) {
	user, err := uc.userRepo.FindOneByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if valid, err := utils.PasswordMatches(user.Password, req.Password); err != nil || !valid {
		return nil, errno.ErrInvalidCredentials
	}

	return nil, nil 
}

func (uc *UserUsecase) UserProfile(ctx context.Context, req *types.ProfileRequest) (*types.ProfileResponse, error) {

	return nil, nil 
}

func (uc *UserUsecase) UserEditInfo(ctx context.Context, req *types.EditInfoRequest) (*types.EditInfoResponse, error) {

	return nil, nil 
}