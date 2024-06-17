package data

import (
	"context"
	"errors"

	"we-backend/internal/biz"
	"we-backend/internal/types"
	"we-backend/pkg/errno"

	"github.com/rs/zerolog/log"
)

type userRepo struct {
	database  UserDatabase
	cache     UserCache 
}

func NewUseRepo(database UserDatabase, cache UserCache) biz.UserRepo {
	return &userRepo{
		database: database,
		cache:    cache,
	}
}

func (impl *userRepo) Insert(ctx context.Context, user types.User) (int64, error) {
	return impl.database.Insert(ctx, user)
}

func (impl *userRepo) FindOne(ctx context.Context, id int64) (*types.User, error) {
	
	/*
	
	缓存里面有数据
	缓存里面没有数据
	缓存出错了，你也不知道有没有数据
	
	 */
	
	user, err := impl.cache.Get(ctx, id)

	switch {
	case errors.Is(err, errno.ErrKeyNoFound):
		// 没数据，那就去数据库里面加载
		user, err = impl.database.FindOne(ctx, id)
		if err != nil {
			return nil, err
		}

		if err := impl.cache.Set(ctx, *user); err != nil {
			// 这里怎么办？打日志，做监控
			log.Info().Err(err).Msg("failed to set user")
		}

		return user, nil 

	case err != nil:
		// 偶发错误，后期新增字段对不上 io.EOF，或者严重错误，10w qps 打崩

		// 这里怎么办，要不要去数据库中加载
		// 1. 选择加载，做好兜底，万一 redis 崩了，要保护好数据库，通过数据库限流
		// 2. 选择不加载，用户体验差一点
		return nil, err 

	default:
		// 必然是有数据
		return user, nil 
	}
}

func (impl *userRepo) FindOneByEmail(ctx context.Context, email string) (*types.User, error) {
	return impl.database.FindOneByEmail(ctx, email) 
}

func (impl *userRepo) FindOneByMobile(ctx context.Context, mobile string) (*types.User, error) {
	return impl.database.FindOneByMobile(ctx, mobile)
}

func (impl *userRepo) Exists(ctx context.Context, id int64) (bool, error) {
	return impl.database.Exists(ctx, id)
}

func (impl *userRepo) Delete(ctx context.Context, id int64) error {
	return impl.database.Delete(ctx, id)
}

func (impl *userRepo) Update(ctx context.Context, user types.User) error {
	return impl.database.Update(ctx, user)
}

func (impl *userRepo)AllUsers(ctx context.Context) ([]*types.User, error) {
	return impl.database.AllUsers(ctx)
}