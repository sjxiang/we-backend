package we


var (
	// 通用
	Success                = NewErrNo(OK, "成功")

	ErrInternal            = NewErrNo(InternalError, "服务器开小差啦，稍后再来试一试")
	ErrMissingParameter    = NewErrNo(MissingParameter, "缺少参数，请重新输入")
	ErrInvalidParameter    = NewErrNo(InvalidParameter, "参数错误，请重新输入")
	
	ErrDatabase            = NewErrNo(DatabaseError, "数据库繁忙，请稍后再试")
	ErrNotFound            = NewErrNo(RecordNoFound, "记录未找到")
	ErrAlreadyExists       = NewErrNo(RecordAlreadyExists, "记录已存在")
	ErrDuplicatedEntry     = NewErrNo(DuplicatedEntry, "重复条目")  
	
	ErrCacheKeyNoFound     = NewErrNo(KeyNoFound, "key 不存在")

	// 用户模块

	// Email is already taken on
	ErrEmailTaken          = NewErrNo(1, "该电子邮件地址已被使用")
	ErrNotLogin            = NewErrNo(NotLogin, "用户未登录，请重新登录")
	ErrInvalidCredentials  = NewErrNo(InvalidCredentials, "用户名或者密码不正确，请重新输入")
	ErrAccessDenied        = NewErrNo(AccessDenied, "用户无权限")

	ErrInvalidToken        = NewErrNo(InvalidToken, "token 无效，请重新登录")
	ErrExpiredToken        = NewErrNo(ExpiredToken, "token 过期，请重新登录")

	ErrOtpSendTooMany      = NewErrNo(OtpSendTooMany, "otp 发送太频繁，请稍后再试")
	ErrOtpVerifyTooMany    = NewErrNo(OtpVerifyTooMany, "otp 验证太频繁，请稍后再试")
)

// 注册，用户已注册
// 登录，用户未注册

