package data

import (
	"context"

	"we-backend/pkg/config"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func NewDB(cfg *config.Config) *gorm.DB {
	
	dsn := cfg.GetMySQLDefaultDSN()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to database")
	}
	
	if cfg.EnableDebugWithSQL() {
		database = database.Debug()
	}

	return database
}

func NewCache(cfg *config.Config) *redis.Client {
	
	addr := cfg.GetRedisAddr()

	cache := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisPassword,  // no password set
		DB:       cfg.RedisDatabase,  // use default DB 0
	})
	_, err := cache.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to cache")
	}

	return cache
}
