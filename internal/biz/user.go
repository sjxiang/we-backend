package biz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"we-backend/pkg/errno"
	"we-backend/pkg/utils"
	"we-backend/internal/service/token"
	"we-backend/internal/types"
	
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo     UserRepo
	tokenService token.TokenService
}

func NewUserUsecase(userRepo UserRepo, tokenService token.TokenService) *UserUsecase {
	return &UserUsecase{userRepo: userRepo, tokenService: tokenService}
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

	rsp := &types.RegisterResponse{
		UserID: id,
	}

	return rsp, nil 
}


func (uc *UserUsecase) UserLogin(ctx context.Context, req *types.LoginRequest) (*types.LoginResponse, error) {
	user, err := uc.userRepo.FindOneByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if valid, err := utils.PasswordMatches(user.Password, req.Password); err != nil || !valid {
		return nil, errno.ErrInvalidCredentials
	}

	fmt.Println(user)
	accessToken, accessPayload, err := uc.tokenService.CreateToken(user.ID, user.Email, time.Duration(144))
	if err != nil {
		return nil, err 
	}

	fmt.Println(accessToken)
	fmt.Println(accessPayload)

	rsp := &types.LoginResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	return rsp, nil 
}

func (uc *UserUsecase) UserProfile(ctx context.Context, req *types.ProfileRequest) (*types.ProfileResponse, error) {
	user, err := uc.userRepo.FindOne(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	rsp := &types.ProfileResponse{
		User: *user,
	}
	fmt.Println(rsp)

	return rsp, nil 
}

func (uc *UserUsecase) UserEditInfo(ctx context.Context, req *types.EditInfoRequest) (*types.EditInfoResponse, error) {

	return nil, nil 
}