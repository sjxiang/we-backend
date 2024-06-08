package errno


// 成功返回
const OK uint32 = 200

/*** 前3位代表业务，后3位代表具体功能 ***/

// 全局错误码
const ServiceUnavailableError uint32 = 100_001
const ParamError uint32= 100_002
const DBError uint32 = 100_003
const RecordNoFoundError uint32 = 100_004
const RecordAlreadyExistsError uint32 = 100_005
const DuplicateEntryError uint32 = 30003


// 用户模块
const (
	InvalidCredentialsError uint32 = 200_001
	AuthorizationFailedError uint32 = 200_002
	TokenExpiredError uint32 = 200_001
)


var (
	// 通用
	Success                = NewErrNo(OK, "成功")

	ErrServiceUnavailable  = NewErrNo(ServiceUnavailableError, "服务器开小差啦，稍后再来试一试")
	ErrParam               = NewErrNo(ParamError, "参数错误")
	ErrDatabase            = NewErrNo(DBError, "数据库繁忙，请稍后再试")

	ErrRecordNoFound       = NewErrNo(RecordNoFoundError, "no matching record found")
	ErrRecordAlreadyExists = NewErrNo(RecordAlreadyExistsError, "record already exists")
	ErrDuplicateEntry      = NewErrNo(DuplicateEntryError, "duplicate entry")  
	

	// 用户模块
	ErrInvalidCredentials  = NewErrNo(InvalidCredentialsError, "用户名或密码不正确，请重新输入")
	ErrTokenExpire         = NewErrNo(TokenExpiredError, "token失效，请重新登录") 
	ErrAuthorizationFailed = NewErrNo(AuthorizationFailedError, "认证失败，请重新登录")
)


// 	ErrNoAuth                 = "请求头中的auth为空"
// 	ErrAuthFormatInvalid      = "请求头中的auth格式有错误"



