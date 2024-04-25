package biz

import (
	"context"
	"errors"
	"fmt"

	"we-backend/pkg/types"
	"we-backend/pkg/utils"
	"we-backend/pkg/x"
)


type UserRepo interface {
	Insert(ctx context.Context, arg types.CreateUserParams) error
	CreateUser(ctx context.Context, arg types.CreateUserParams) error
	FindUserByEmail(ctx context.Context, email string) (*types.User, error)  
	FindUserByMobile(ctx context.Context, mobile string) (types.User, error)
}

type UserUsecase struct {
	ur UserRepo
}

func NewUserUsecase(ur UserRepo) *UserUsecase {
	return &UserUsecase{ur: ur}
}

func (impl *UserUsecase) UserSignup(ctx context.Context, email, password string) error {
	
	hashedPassword, err := utils.GenerateHashFromPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash the password\n %w", err)
	}

	arg := types.CreateUserParams{
		Email:    email,
		Password: hashedPassword,
	}
	if err := impl.ur.CreateUser(ctx, arg); err != nil {
		
		if errors.Is(err, x.ErrDuplicateEmail) {
			return x.ErrUserAlreadyExists
		}
    
		return err
	}

	return nil 
}


func (impl *UserUsecase) UserLogin(ctx context.Context, email, password string) (*types.User, error) {

	user, err := impl.ur.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, x.ErrNoRecord) {
			return nil, x.ErrInvalidCredentials
		}

		return nil, err  // "数据库繁忙，请稍后再试"
	}

	match, err := utils.PasswordMatches(user.Password, password)
	if err != nil {
		return nil, fmt.Errorf("failed to match hash and password\n %w", err)  // "服务器开小差啦，稍后再来试一试"
	}
	if !match {
		return nil, x.ErrInvalidCredentials
	}
	
	return user, nil 
}
