package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"we-backend/internal/conf"
	"we-backend/internal/types"
	"we-backend/pkg/we"
)


type userCacheImpl struct {
	cmd        redis.Cmdable  
	expiration time.Duration
}

func NewUserCache(cache *redis.Client, cfg *conf.Config) UserCache {
	return &userCacheImpl{
		cmd:        cache,
		expiration: time.Minute * time.Duration(cfg.RedisExpiration),
	}
}
 
func (impl *userCacheImpl) Get(ctx context.Context, id int64) (types.User, error) {
	
	key := userIDKey(id)

	value, err := impl.cmd.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return types.User{}, we.ErrCacheKeyNoFound
	} else if err != nil {
		return types.User{}, fmt.Errorf("get user: %w", err)
	}

	var user types.User
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return types.User{}, fmt.Errorf("failed to decode user json: %w", err)
	}

	return user, nil 
}

// 使用 Redis Set 创建/覆盖缓存，并设置 1 分钟 TTL 过期时间 
func (impl *userCacheImpl) Set(ctx context.Context, user types.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to encode user: %w", err)
	}

	key := userIDKey(user.ID)
	
	return impl.cmd.Set(ctx, key, data, impl.expiration).Err()
}

func (impl *userCacheImpl) Del(ctx context.Context, id int64) error {
	key := userIDKey(id)
	
	err := impl.cmd.Del(ctx, key).Err()
	switch {
	case errors.Is(err, redis.Nil):
		return we.ErrCacheKeyNoFound
	case err != nil:
		return fmt.Errorf("failed to remove user: %w", err)
	default:
		return nil 
	}
}

func userIDKey(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
