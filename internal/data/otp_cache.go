package data

import (
	"context"
	_ "embed"
	"fmt"
	"errors"

	"github.com/redis/go-redis/v9"

	"we-backend/pkg/errno"
)

//go:embed lua/set_code.lua
var luaSetCode string
//go:embed lua/verify_code.lua
var luaVerifyCode string

type otpCacheImpl struct {
	cmd        redis.Cmdable	
}

func NewOtpCache(client redis.Cmdable) OtpCache {
	return &otpCacheImpl{
		cmd: client,
	}
}

func (impl *otpCacheImpl) Set(ctx context.Context, biz, phoneNumber, code string) error {
	result, err := impl.cmd.Eval(ctx, luaSetCode, []string{otpKey(biz, phoneNumber)}, code).Int()
	// 打印日志
	if err != nil {
		// 调用 redis 出了问题
		return err
	}
	
	switch result {
	case -2:
		// 系统异常
		return errors.New("验证码存在，但是没有过期时间")
	case -1:
		// 发送太频繁
		return errno.ErrOtpSendTooMany
	default:
		// 正常
		return nil  
	}
}

func (impl *otpCacheImpl) Verify(ctx context.Context, biz, phoneNumber, inputCode string) (bool, error) {
	result, err := impl.cmd.Eval(ctx, luaVerifyCode, []string{otpKey(biz, phoneNumber)}, inputCode).Int()
	if err != nil {
		return false, err
	}

	switch result {
	case -2:
		// 用户手一抖 输入错
		return false, nil 
	case -1:
		// 用户一直输错 频繁出现这个错误 就要告警 合理怀疑 有人搞你
		return false, errno.ErrOtpVerifyTooMany
	default:
		// 输入对了
		return true, nil 
	}
}

func otpKey(biz, phoneNumber string) string {
	// otp:login:188xxxxoooo
	return fmt.Sprintf("otp:%s:%s", biz, phoneNumber)
}
