package limiter

import (
	"context"
	"time"
	"we-backend/internal/conf"

	"github.com/redis/go-redis/v9"
)

type RateLimitService interface {
	Limit(ctx context.Context, key string) (bool, error)  // 是否触发限流
}

func NewRateLimitService(cmd redis.Cmdable, cfg *conf.Config) RateLimitService {
	return &wrapper{
		cmd:      cmd,
		interval: time.Duration(cfg.LimitInternal),  // 比方一分钟上限 100 个请求
		rate:     cfg.LimitRate,
	}
}