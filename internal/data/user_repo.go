package data

import (
	"context"
	"errors"

	"we-backend/internal/biz"
	"we-backend/internal/types"
	"we-backend/pkg/we"

	"github.com/rs/zerolog/log"
)

type userRepo struct {
	database  UserDatabase
	cache     UserCache 
}

func NewUserRepo(database UserDatabase, cache UserCache) biz.UserRepo {
	return &userRepo{
		database: database,
		cache:    cache,
	}
}

func (impl *userRepo) Create(ctx context.Context, user types.User) (types.User, error) {
	return impl.database.Create(ctx, user)
}
	
func (impl *userRepo) GetByID(ctx context.Context, id int64) (types.User, error) {

	/*
	
	缓存里面有数据
	缓存里面没有数据
	缓存出错了，你也不知道有没有数据
	
	 */
	
	u, err := impl.cache.Get(ctx, id)

	switch {
	case errors.Is(err, we.ErrCacheKeyNoFound):
		// 没数据，那就去数据库里面加载
		user, err := impl.database.GetByID(ctx, id)
		if err != nil {
			return types.User{}, err
		}

		if err := impl.cache.Set(ctx, user); err != nil {
			// 这里怎么办？打日志，做监控
			log.Info().Err(err).Msg("failed to set user")
		}

		return user, nil 

	case err != nil:
		// 偶发错误，后期新增字段对不上 io.EOF，或者严重错误，10w qps 打崩

		// 这里怎么办，要不要去数据库中加载
		// 1. 选择加载，做好兜底，万一 redis 崩了，要保护好数据库，通过数据库限流
		// 2. 选择不加载，用户体验差一点
		return types.User{}, err 

	default:
		// 必然是有数据
		return u, nil 
	}
}

func (impl *userRepo) GetByEmail(ctx context.Context, email string) (types.User, error) {
	return impl.database.GetByEmail(ctx, email)
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

func (impl *userRepo)All(ctx context.Context) ([]types.User, error) {
	return impl.database.All(ctx)
}