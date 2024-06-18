package data

import (
	"context"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"we-backend/internal/types"
	"we-backend/pkg/errno"
)


type userDatabaseImpl struct {
	db *gorm.DB 
}

func NewUserDatabase(db *gorm.DB) UserDatabase {
	return &userDatabaseImpl{
		db: db,
	}
}

func (impl *userDatabaseImpl) Insert(ctx context.Context, user types.User) (int64, error) {
	
	err := impl.db.Table("users").Create(&user).Error
	
	switch {
	case err != nil:
		var mysqlError *mysql.MySQLError
		
		if errors.As(err, &mysqlError) {
			switch {
			case mysqlError.Number == 1062 && 
					strings.Contains(mysqlError.Message, "users.uk_email"):
				return 0, errno.ErrDuplicatedEntry.WithMessage("邮箱重复")
			default:
				return 0, errno.ErrDuplicatedEntry
			}
		}

		return 0, err
	default:
		return user.ID, nil 
	}
}


func (impl *userDatabaseImpl) FindOne(ctx context.Context, id int64) (*types.User, error) {
	var item types.User

	err := impl.db.Table("users").Where("id = ?", id).First(&item).Error
	
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errno.ErrRecordNoFound
	case err != nil:
		return nil, err
	default:
		return &item, nil 
	} 
}


func (impl *userDatabaseImpl) FindOneByEmail(ctx context.Context, email string) (*types.User, error) {
	var row types.User

	err := impl.db.Table("users").Where("email = ?", email).First(&row).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return nil, errno.ErrRecordNoFound
	case nil:
		return &row, nil
	default:
		return nil, err 
	} 
}

func (impl *userDatabaseImpl) FindOneByMobile(ctx context.Context, mobile string) (*types.User, error) {
	var resp types.User
	
	err := impl.db.Table("users").Where("mobile = ?", mobile).First(&resp).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errno.ErrRecordNoFound
	case err != nil:
		return nil, err 
	default:
		return &resp, err 
	} 
}

func (impl *userDatabaseImpl) Exists(ctx context.Context, id int64) (bool, error) {
	var exists bool

	stmt := "SELECT EXISTS(SELECT true FROM user WHERE id = ?)"

	err := impl.db.Raw(stmt, id).Scan(&exists).Error

	return exists, err
}

func (impl *userDatabaseImpl) Delete(ctx context.Context, id int64) error {
	
	if err := impl.db.Transaction(func(tx *gorm.DB) error {
		var exists bool

		stmt := "SELECT EXISTS(SELECT true FROM user WHERE id = ?)"

		if err := tx.Raw(stmt, id).Scan(&exists).Error; err != nil {
			return err
		}
		if !exists {
			return errno.ErrRecordNoFound
		}
		
		if err := tx.Delete(&types.User{}, id).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err 
	}

	return nil
}




func (impl *userDatabaseImpl) Update(ctx context.Context, user types.User) error {
	// 这种写法依赖于 GORM 的零值和主键更新特性
	// Update 非零值 WHERE id = ?

	return impl.db.Table("users").Where("id = ?", user.ID).
		Updates(map[string]any{
			"nickname": user.Nickname,
			"avatar":   user.Avatar,
			"intro":    user.Avatar,
			"birthday": user.Birthday,
	}).Error
}

func (impl *userDatabaseImpl) AllUsers(ctx context.Context) ([]*types.User, error) {
	return nil, nil
}

func (impl *userDatabaseImpl) ResetPassword(ctx context.Context, id int64, password string) error {	
	return nil 
}

