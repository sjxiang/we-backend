package data

import (
	"context"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"we-backend/internal/conf"
	"we-backend/internal/types"
	"we-backend/pkg/faker"
)

func Test_raw_insert_user(t *testing.T) {
	cfg, _ := conf.LoadConfig()

	db := NewRawDatabase(cfg)
	
	repo := &userRawDatabase{rawDB: db}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)
	id , err := repo.Insert(context.TODO(), types.User{
		Nickname: faker.Username(),
		Email:    faker.Email(),
		Mobile:   "1330000"+faker.RandIntSpec(),
		Password: string(hashedPassword),
		Avatar:   "jisoo.jpeg",
		Intro:    "life is fucking movie.",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("用户 id %d\n", id)
}


func Test_raw_delete_user(t *testing.T) {
	cfg, _ := conf.LoadConfig()

	db := NewRawDatabase(cfg)
	
	repo := &userRawDatabase{rawDB: db}

	if err := repo.Delete(context.TODO(), 25); err != nil {
		t.Fatal(err)
	}

}
