package data

import (
	"context"
	"errors"
	"strings"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"we-backend/internal/types"
	"we-backend/pkg/we"
)


type userDatabaseImpl struct {
	DB *gorm.DB 
}

func NewUserDatabase(db *gorm.DB) UserDatabase {
	return &userDatabaseImpl{
		DB: db,
	}
}

func (impl *userDatabaseImpl) Create(ctx context.Context, user types.User) (types.User, error) {
	
	if err := impl.DB.Table("users").Create(&user).Error; err != nil {
		
		var mysqlError *mysql.MySQLError
		
		if errors.As(err, &mysqlError) {
			switch {
			case mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users.uk_email"):
				return types.User{}, we.ErrAlreadyExists  
			default:
				return types.User{}, we.ErrAlreadyExists
			}
		}

		return types.User{}, fmt.Errorf("error insert: %v", err)
	}

	return user, nil 
}

func (impl *userDatabaseImpl) GetByEmail(ctx context.Context, email string) (types.User, error) {
	
	u := types.User{}

	if err := impl.DB.Table("users").Where("email = ?", email).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return types.User{}, we.ErrNotFound
		}

		return types.User{}, fmt.Errorf("error select: %v", err)
	}

	return u, nil 
}

func (impl *userDatabaseImpl) GetByID(ctx context.Context, id int64) (types.User, error) {
	
	u := types.User{}

	err := impl.DB.Table("users").Where("id = ?", id).First(&u).Error
	
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return types.User{}, we.ErrNotFound
	case err != nil:
		return types.User{}, fmt.Errorf("error select: %v", err)
	default:
		return u, nil 
	}  
}

func (impl *userDatabaseImpl) Exists(ctx context.Context, id int64) (bool, error) {
	var exists bool

	stmt := "SELECT EXISTS(SELECT true FROM user WHERE id = ?)"

	err := impl.DB.Raw(stmt, id).Scan(&exists).Error

	return exists, err
}

func (impl *userDatabaseImpl) Delete(ctx context.Context, id int64) error {
	
	if err := impl.DB.Transaction(func(tx *gorm.DB) error {
		return deleteUser(tx, id)
	}); err != nil {
		return err 
	}

	return nil
}


func deleteUser(tx *gorm.DB, id int64) error {
	var exists bool

	stmt := "SELECT EXISTS(SELECT true FROM user WHERE id = ?)"

	if err := tx.Raw(stmt, id).Scan(&exists).Error; err != nil {
		return err
	}
	if !exists {
		return we.ErrNotFound
	}
	
	if err := tx.Where("id = ?", id).Delete(&types.User{}).Error; err != nil {
		return err
	}

	return nil
}


func (impl *userDatabaseImpl) Update(ctx context.Context, user types.User) error {
	// 这种写法依赖于 GORM 的零值和主键更新特性
	// Update 非零值 WHERE id = ?

	return impl.DB.Table("users").Where("id = ?", user.ID).
		Updates(map[string]any{
			"nickname": user.Nickname,
			"avatar":   user.Avatar,
			"intro":    user.Avatar,
			"birthday": user.Birthday,
	}).Error
}

func (impl *userDatabaseImpl) All(ctx context.Context) ([]types.User, error) {
	
	var uu []types.User
	
	if err := impl.DB.Table("users").Find(&uu).Error; err != nil {
		return nil, err
	}
	
	return uu, nil 
}

func (impl *userDatabaseImpl) ResetPassword(ctx context.Context, id int64, password string) error {	
	return nil 
}

