package we

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

const KeyNoFound          uint32 = 100_300


// 用户模块
const (
	NotLogin            uint32 = 200_100
	InvalidCredentials  uint32 = 200_101
	AccessDenied        uint32 = 200_102

	InvalidToken        uint32 = 200_200
	ExpiredToken        uint32 = 200_201
)

// otp 模块
const (
	OtpSendTooMany      uint32 = 300_100
	OtpVerifyTooMany    uint32 = 300_101
)

