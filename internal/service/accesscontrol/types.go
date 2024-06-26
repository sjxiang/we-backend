package accesscontrol

import (
	"context"
	"time"
	"we-backend/internal/conf"

	"github.com/redis/go-redis/v9"
)

// 接入控制（限流）
type AccessControlService interface {
	Limit(ctx context.Context, key string) (bool, error)  // 是否触发限流
}

func NewAccessControlService(cmd redis.Cmdable, cfg *conf.Config) AccessControlService {
	return &wrapper{
		cmd:      cmd,
		interval: time.Duration(cfg.LimitInternal),  // 比方一分钟上限 10 个请求
		rate:     cfg.LimitRate,
	}
}