package data

import (
	"context"
	_ "embed"
	"fmt"
	"errors"

	"github.com/redis/go-redis/v9"

	"we-backend/pkg/we"
	"we-backend/internal/biz"
)

type captchaRepo struct {
	cmd   redis.Cmdable	
}

//go:embed lua/set_code.lua
var luaSetCode string
//go:embed lua/verify_code.lua
var luaVerifyCode string


func NewCaptchaRepo(client redis.Cmdable) biz.CaptchaRepo {
	return &captchaRepo{	
		cmd: client,
	}
}

func captchaKey(biz, phoneNumber string) string {
	// captcha:login:188xxxxoooo
	return fmt.Sprintf("otp:%s:%s", biz, phoneNumber)
}


func (impl *captchaRepo) Store(ctx context.Context, biz, phoneNumber, code string) error {
	key := captchaKey(biz, phoneNumber)

	result, err := impl.cmd.Eval(ctx, luaSetCode, []string{key}, code).Int()
	if err != nil {
		// 打印日志，调用 redis 出了问题
		return err
	}
	
	switch result {
	case -2:
		// 系统异常
		return errors.New("验证码存在，但是没有过期时间")
	case -1:
		// 发送太频繁
		return we.ErrOtpSendTooMany
	default:
		// 正常
		return nil  
	}
}

func (impl *captchaRepo) Verify(ctx context.Context, biz, phoneNumber, inputCode string) (bool, error) {
	key := captchaKey(biz, phoneNumber)

	result, err := impl.cmd.Eval(ctx, luaVerifyCode, []string{key}, inputCode).Int()
	if err != nil {
		return false, err
	}

	switch result {
	case -2:
		// 用户手一抖 输入错误
		return false, nil 
	case -1:
		// 用户一直输错，频繁出现这个错误，就要告警，合理怀疑，有人搞你
		return false, we.ErrOtpVerifyTooMany
	default:
		// 输入对了
		return true, nil 
	} 
}