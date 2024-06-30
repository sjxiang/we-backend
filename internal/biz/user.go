package biz

import (
	"context"
	
	"we-backend/internal/types"
)

type UserUsecase struct {
	UserRepo     UserRepo
}

func NewUserUsecase(ur UserRepo) *UserUsecase {
	return &UserUsecase{
		UserRepo: ur, 
	}
}


func (uc *UserUsecase) Me(ctx context.Context, input types.MeInput) (types.MeResponse, error) {
	
	user, err := uc.UserRepo.GetByID(ctx, input.UserID)
	if err != nil {
		return types.MeResponse{}, err
	}
	
	return types.MeResponse{
		User: mapUser(user),
	}, nil 
}

// 函数名 ExportUserForFeedback
func mapUser(u types.User) types.User {
	return types.User{
		Nickname:  u.Nickname,
		Mobile:    u.Mobile,
		Email:     u.Email,
		Intro:     u.Intro,
		Avatar:    u.Avatar,
		Birthday:  u.Birthday,
		CreatedAt: u.CreatedAt,
	}
}


func (uc *UserUsecase) Edit(ctx context.Context, input types.EditInput) error {
	
	newUser := types.User{
		ID:       input.UserID,
		Nickname: input.Nickname,
		Avatar:   input.Avatar,
		Intro:    input.Intro,
		Birthday: input.Birthday, 
	}

	return uc.UserRepo.Update(ctx, newUser)
}


func (uc *UserUsecase) All(ctx context.Context) (types.AllResponse, error) {

	uu, err := uc.UserRepo.All(ctx)
	if err != nil {
		return types.AllResponse{}, err
	}

	return types.AllResponse{
		Users: mapUsers(uu),
	}, nil 
}

func convertUser(u types.User) *types.User {
	return &types.User{
		ID:        u.ID,
		Nickname:  u.Nickname,
		Mobile:    u.Mobile,
		Email:     u.Email,
		Intro:     u.Intro,
		Avatar:    u.Avatar,
		CreatedAt: u.CreatedAt,
	}
}

func mapUsers(uu []types.User) []*types.User {
	users := make([]*types.User, len(uu))

	for i, u := range uu {
		users[i] = convertUser(u)	
	}

	return users
}
