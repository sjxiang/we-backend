package consts


// 常量

const (

	// config
	SERVER_MODE_DEV                = "dev"  // 开发环境
	SERVER_MODE_PRO                      = "pro"  // 生产环境
	ServerAuthzDeployModeSingle          = "single"  // 单实例，memory
	ServerAuthzDeployModeMulti           = "multi"   // 多实例，redis
	
	// data
	TABLE_NAME_USER    = "user"
	TABLE_NAME_ARTICLE = "article"
	
	// key
	SessionKeyUserId    = "user_id"
	SessionKeyLastTime  = "last_time"

	 
)