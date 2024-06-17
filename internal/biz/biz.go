package biz

import (
	"context"

	"we-backend/internal/types"
)


type UserRepo interface {
	Insert(ctx context.Context, user types.User) (int64, error)
	FindOne(ctx context.Context, id int64) (*types.User, error)
	FindOneByEmail(ctx context.Context, email string) (*types.User, error) 
	FindOneByMobile(ctx context.Context, mobile string) (*types.User, error) 
	Exists(ctx context.Context, id int64) (bool, error) 
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user types.User) error 
	AllUsers(ctx context.Context) ([]*types.User, error)
}


type UserCache interface {
	Get(ctx context.Context, id int64) (*types.User, error)
	Set(ctx context.Context, user types.User) error
	Del(ctx context.Context, id int64) error
}
