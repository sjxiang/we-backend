package config

import (
	"fmt"
	"we-backend/pkg/consts"
)

type Config struct {
	ServerHost             string
	ServerPort             string
	ServerMode             string 
	SecretKey              string
	ServerAuthzMode        int8

	MySQLHost              string 
	MySQLPort              string 
	MySQLUser              string 
	MySQLPassword          string 
	MySQLDatabase          string 
	
	RedisNetworkType       string
	RedisHost              string 
	RedisPort              string 
	RedisPassword          string 
	RedisDatabase          int 
}


func LoadConfig() (config *Config, err error) {
	cfg := &Config{}
	
	cfg.MySQLHost         = "127.0.0.1"
	cfg.MySQLPort         = "8001"
	cfg.ServerMode        = consts.SERVER_MODE_DEV
	cfg.SecretKey         = "8xEMrWkBARcDDYQ"
	cfg.ServerAuthzMode   = consts.ServerAuthzModeCookie

	cfg.MySQLHost         = "127.0.0.1"
	cfg.MySQLPort         = "13306"
	cfg.MySQLUser         = "root"
	cfg.MySQLPassword     = "my-secret-pw"
	cfg.MySQLDatabase     = "we_backend"

	cfg.RedisNetworkType  = "tcp"
	cfg.RedisHost         = "127.0.0.1"
	cfg.RedisPort         = "16379"
	cfg.RedisPassword     = ""
	cfg.RedisDatabase     = 0
	
	return cfg, nil
}

func (cfg *Config) GetServerAddr() string {
	// "127.0.0.1:8001"
	return fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)
}

// 是否启用 SQL 调试
func (cfg *Config) EnableDebugWithSQL() bool {
	return cfg.ServerMode == consts.SERVER_MODE_DEV
}

func (cfg *Config) GetSecretKey() string {
	return cfg.SecretKey
}

func (cfg *Config) GetMySQLDefaultDSN() string {
	// "root:my-secret-pw@tcp(127.0.0.1:13306)/we_backend?charset=utf8&parseTime=True&loc=Local")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
		cfg.MySQLUser, 
		cfg.MySQLPassword, 
		cfg.MySQLHost, 
		cfg.MySQLPort, 
		cfg.MySQLDatabase)
}


func (cfg *Config) GetRedisAddr() string {
	// "localhost:6379"
	return fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
}

func (cfg *Config) EnableAuthzCookie() bool {
	return cfg.ServerAuthzMode == consts.ServerAuthzModeCookie
}

func (cfg *Config) EnableAuthzMultiSession() bool {
	return cfg.ServerAuthzMode == consts.ServerAuthzModeSessionMulti
}

func (cfg *Config) EnableAuthzSingleSession() bool {
	return cfg.ServerAuthzMode == consts.ServerAuthzModeSessionSingle
}

func (cfg *Config) EnableAuthzJWT() bool {
	return cfg.ServerAuthzMode == consts.ServerAuthzModeJWT
}