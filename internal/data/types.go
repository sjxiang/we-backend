package data

import (
	"context"

	"we-backend/internal/types"
)

type UserDatabase interface {
	Create(ctx context.Context, user types.User) (types.User, error)
	GetByEmail(ctx context.Context, email string) (types.User, error)
	GetByID(ctx context.Context, id int64) (types.User, error)
	Exists(ctx context.Context, id int64) (bool, error) 
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, user types.User) error 
	All(ctx context.Context) ([]types.User, error)
}


type UserCache interface {
	Get(ctx context.Context, id int64) (types.User, error)
	Set(ctx context.Context, user types.User) error
	Del(ctx context.Context, id int64) error
}

