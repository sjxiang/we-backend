package errno

// 成功返回
const OK uint32 = 200

/*** 前3位代表业务，后3位代表具体功能 ***/

// 全局错误码
const InternalError    uint32 = 100_100  // 异常
const InvalidParameter uint32 = 100_101
const MissingParameter uint32 = 100_102

const DatabaseError       uint32 = 100_200	
const RecordNoFound       uint32 = 100_201
const RecordAlreadyExists uint32 = 100_202
const DuplicatedEntry     uint32 = 100_203

// 用户模块
const (
	NotLogin            uint32 = 200_100
	InvalidCredentials  uint32 = 200_101
	AuthorizationFailed uint32 = 200_102
	AccessDenied        uint32 = 200_103

	InvalidToken        uint32 = 200_200
	ExpiredToken        uint32 = 200_201
)


var (
	// 通用
	Success                = NewErrNo(OK, "成功")

	ErrInternal            = NewErrNo(InternalError, "服务器开小差啦，稍后再来试一试")
	ErrInvalidParameter    = NewErrNo(InvalidParameter, "参数错误")
	ErrMissingParameter    = NewErrNo(MissingParameter, "缺少参数")
	
	ErrDatabase            = NewErrNo(DatabaseError, "数据库繁忙，请稍后再试")
	ErrRecordNoFound       = NewErrNo(RecordNoFound, "资源不存在")
	ErrRecordAlreadyExists = NewErrNo(RecordAlreadyExists, "资源已存在")
	ErrDuplicatedEntry      = NewErrNo(DuplicatedEntry, "重复条目")  
	

	// 用户模块
	ErrNotLogin            = NewErrNo(NotLogin, "未登录")
	ErrInvalidCredentials  = NewErrNo(InvalidCredentials, "用户名或密码不正确，请重新输入")
	ErrAuthorizationFailed = NewErrNo(AuthorizationFailed, "JWT 认证失败，请稍后重试")
	ErrAccessDenied        = NewErrNo(AccessDenied, "无权限")

	ErrInvalidToken        = NewErrNo(InvalidToken, "token 无效")
	ErrExpiredToken        = NewErrNo(ExpiredToken, "token 过期，请重新登录")
)


// 注册，用户已注册
// 登录，用户未注册