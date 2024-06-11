package data

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"we-backend/internal/conf"
)


func NewDB(cfg *conf.Config) *gorm.DB {
	
	dsn := cfg.GetMySQLDefaultDSN()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to database")
	}
	
	if cfg.EnableDebugWithSQL {
		database = database.Debug()
	}

	// setup sql connection pool
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to sql connection pool")
	}
	
	sqlDB.SetMaxIdleConns(cfg.MySQLMaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MySQLMaxOpenConns)

	return database
}

func NewCache(cfg *conf.Config) *redis.Client {
	
	addr := cfg.GetRedisAddr()

	cache := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisPassword,  // no password set
		DB:       cfg.RedisDB,        // use default DB 0
	})
	_, err := cache.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to cache")
	}

	return cache
}



// mysql <raw sql>
func NewRawDatabase(cfg *conf.Config) *sql.DB {
	connection, err := openDB(cfg.GetMySQLDefaultDSN())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to raw database")
	}

	return connection
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}