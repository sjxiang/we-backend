package data

import (
	"context"
	"testing"
	"we-backend/pkg/config"
	"we-backend/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

func Test_raw_insert_user(t *testing.T) {
	cfg, _ := config.LoadConfig()

	db := NewRawDatabase(cfg)
	
	repo := &userRawDatabase{rawDB: db}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)
	id , err := repo.Insert(context.TODO(), types.User{
		Nickname: "admin",
		Email:    "admin123@qq.com",
		Mobile:   "13300001111",
		Password: string(hashedPassword),
		Avatar:   "jisoo.jpeg",
		Intro:    "life is fucking movie.",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("用户 id %d\n", id)
}

func Test_raw_find_one(t *testing.T) {
	
}

func Test_raw_find_one_by_email(t *testing.T) {
	
}

func Test_raw_delete_user(t *testing.T) {
	cfg, _ := config.LoadConfig()

	db := NewRawDatabase(cfg)
	
	repo := &userRawDatabase{rawDB: db}

	if err := repo.Delete(context.TODO(), 25); err != nil {
		t.Fatal(err)
	}

}

func Test_raw_reset_password(t *testing.T) {
	
}

func Test_raw_exists(t *testing.T) {
	
}

func Test_all_users(t *testing.T) {
	cfg, _ := config.LoadConfig()

	db := NewRawDatabase(cfg)
	
	repo := &userRawDatabase{rawDB: db}

	users, err := repo.AllUsers(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(users)
}