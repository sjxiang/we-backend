package consts


// 常量

const (

	// config
	SERVER_MODE_DEV                      = "dev"  // 开发环境
	SERVER_MODE_PRO                      = "pro"  // 生产环境

	SERVER_AUTHZ_MODE_JWT                = 1
	SERVER_AUTHZ_MODE_COOKIE_AND_SESSION = 2
	
	// data
	TABLE_NAME_USER    = "user"
	TABLE_NAME_ARTICLE = "article"
	
	// key
	SessionKeyUserId    = "user_id"
	SessionKeyLastTime  = "last_time"

	 
)