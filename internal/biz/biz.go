package biz

import (
	"context"

	"we-backend/internal/types"
)


type UserRepo interface {
	Create(ctx context.Context, user types.User) (types.User, error)
	GetByEmail(ctx context.Context, email string) (types.User, error)
	GetByID(ctx context.Context, id int64) (types.User, error)
	Exists(ctx context.Context, id int64) (bool, error) 
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user types.User) error 
	All(ctx context.Context) ([]types.User, error)
}


type CaptchaRepo interface {
	Store(ctx context.Context, biz, phoneNumber, code string) error
	Verify(ctx context.Context, biz, phoneNumber, inputCode string) (bool, error)
}