package biz

import "we-backend/pkg/errno"

const (
	InvalidCredentials uint32 = 300_001
)

var (
	ErrInvalidCredentials  = errno.NewErrNo(InvalidCredentials, "用户名或密码不正确，请重新输入")
)

// 注册，用户已注册
// 登录，用户未注册
