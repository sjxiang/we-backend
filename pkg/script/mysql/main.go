package main

import (
	"we-backend/pkg/data"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:13306)/we_backend?charset=utf8&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to database")
	}

	database.AutoMigrate(&data.UserM{})
}
