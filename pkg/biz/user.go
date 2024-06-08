package biz

import (
	"context"

	"we-backend/pkg/types"
)

type UserUsecase struct {
	userRepo UserRepo
}

func NewUserUsecase(userRepo UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uc *UserUsecase) UserRegister(ctx context.Context, req *types.RegisterRequest) (*types.RegisterResponse, error) {

	return nil, nil 
}


func (uc *UserUsecase) UserLogin(ctx context.Context, req *types.LoginRequest) (*types.LoginResponse, error) {

	return nil, nil 
}

func (uc *UserUsecase) UserProfile(ctx context.Context, req *types.ProfileRequest) (*types.ProfileResponse, error) {

	return nil, nil 
}

func (uc *UserUsecase) UserEditInfo(ctx context.Context, req *types.EditInfoRequest) (*types.EditInfoResponse, error) {

	return nil, nil 
}

	
// 	hashedPassword, err := utils.GenerateHashFromPassword(password)
// 	if err != nil {
// 		return fmt.Errorf("failed to hash the password\n %w", err)
// 	}

// 	arg := types.CreateUserParams{
// 		Email:    email,
// 		Password: hashedPassword,
// 	}
// 	if err := impl.ur.CreateUser(ctx, arg); err != nil {
		
// 		if errors.Is(err, x.ErrDuplicateEmail) {
// 			return x.ErrUserAlreadyExists
// 		}
    
// 		return err
// 	}

// 	return nil 
// }


// func (impl *UserUsecase) UserLogin(ctx context.Context, email, password string) (*types.User, error) {

// 	user, err := impl.ur.FindUserByEmail(ctx, email)
// 	if err != nil {
// 		if errors.Is(err, x.ErrNoRecord) {
// 			return nil, x.ErrInvalidCredentials
// 		}

// 		return nil, err  // "数据库繁忙，请稍后再试"
// 	}

// 	match, err := utils.PasswordMatches(user.Password, password)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to match hash and password\n %w", err)  // "服务器开小差啦，稍后再来试一试"
// 	}
// 	if !match {
// 		return nil, x.ErrInvalidCredentials
// 	}
	
// 	return user, nil 
// }
