package access

import (
	"context"
	_ "embed"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:embed slide_window.lua
var luaScript string

type wrapper struct {
	cmd      redis.Cmdable
	interval time.Duration
	rate     int64  // 阈值
}

func (w *wrapper) Limit(ctx context.Context, key string) (bool, error) {
	return w.cmd.Eval(
		ctx, 
		luaScript, 
		[]string{key},
		w.interval.Milliseconds(),  // arg 1
		w.rate,                     // arg 2
		time.Now().UnixMilli(),     // arg 3
	).Bool()
}

// 让 redis 统计请求数量