package data

import (
	"context"
	"testing"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func Test_find_one_by_mobile(t *testing.T) {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:13306)/we_backend?charset=utf8&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to database")
	}

	userRepo := NewUserRepo(database.Debug())
	resp, err := userRepo.FindOneByMobile(context.TODO(), "18812347777")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func Test_find_one_by_email(t *testing.T) {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:13306)/we_backend?charset=utf8&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to database")
	}

	userRepo := NewUserRepo(database.Debug())
	resp, err := userRepo.FindOneByEmail(context.TODO(), "1535484943@qq.com")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}