package x

import (
	"errors"
)


// 统一错误状态码

var (
	
	// ErrEmpty
	// ErrInvalid

	// data
	ErrNoRecord           = errors.New("data: no matching record found")

	ErrDuplicateEmail     = errors.New("data: duplicate email 邮箱冲突")
	ErrDuplicateMobile    = errors.New("data: duplicate mobile 手机号码冲突")
	ErrDuplicateUserID    = errors.New("data: duplicate user id 用户编号冲突")

	
	// biz
	ErrUserAlreadyExists  = errors.New("biz: user already exist 用户已存在")
	ErrUserNotFound       = errors.New("biz: user not found 用户不存在")
	ErrInvalidCredentials = errors.New("biz: invalid credentials 账号或密码不对")
)
