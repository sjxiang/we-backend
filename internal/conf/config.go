package conf

import (
	"fmt"
)

const (
	SERVER_ENV_RELEASE = "release"
	SERVER_ENV_DEV     = "dev"	
)

type Config struct {
	ServerHost             string
	ServerPort             int
	Env                    string
	SecretKey              string
	
	MySQLUser              string 
	MySQLPassword          string 
	MySQLHost              string 
	MySQLPort              int 
	MySQLDatabaseName      string 
	MySQLMaxIdleConns      int
	MySQLMaxOpenConns      int   
	EnableDebugWithSQL     bool  // 是否启用 SQL 调试
	

	RedisHost              string 
	RedisPort              string 
	RedisPassword          string 
	RedisDB                int 
	RedisExpiration        int

	JWTIssuer              string 
	JWTExpirationTime      int64  

	LimitInternal          int64  // 一秒内上限 100 个请求
	LimitRate              int64  
}


func LoadConfig() (config *Config, err error) {
	cfg := &Config{}
	
	cfg.ServerHost        = "127.0.0.1"
	cfg.ServerPort        = 5678
	cfg.SecretKey         = "8xEMrWkBARcDDYQ"
	cfg.Env               = SERVER_ENV_DEV

	cfg.MySQLHost         = "127.0.0.1"
	cfg.MySQLPort         = 3306
	cfg.MySQLUser         = "root"
	cfg.MySQLPassword     = "my-secret-pw"
	cfg.MySQLDatabaseName = "we_backend"

	cfg.RedisHost         = "127.0.0.1"
	cfg.RedisPort         = "6379"
	cfg.RedisPassword     = ""
	cfg.RedisDB           = 0
	cfg.RedisExpiration   = 24
	
	cfg.JWTIssuer         = "we"
	cfg.JWTExpirationTime = 30 * 86400
	
	cfg.LimitInternal     = 60
	cfg.LimitRate         = 5
	
	return cfg, nil
}

func (cfg *Config) GetServerAddr() string {
	// "127.0.0.1:8001"
	return fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort)
}


func (cfg *Config) GetSecretKey() string {
	return cfg.SecretKey
}

func (cfg *Config) GetMySQLDefaultDSN() string {
	// "root:my-secret-pw@tcp(127.0.0.1:13306)/we_backend?charset=utf8&parseTime=True&loc=Local")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", 
		cfg.MySQLUser, 
		cfg.MySQLPassword, 
		cfg.MySQLHost, 
		cfg.MySQLPort, 
		cfg.MySQLDatabaseName)
}

func (cfg *Config) GetRedisAddr() string {
	// "localhost:16379"
	return fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
}

func (cfg *Config) IsLocal() bool {
	return cfg.Env == SERVER_ENV_DEV
}
