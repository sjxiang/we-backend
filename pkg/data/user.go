package data

import (
	"context"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"we-backend/pkg/biz"
	"we-backend/pkg/types"
)


type userDatabase struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) biz.UserRepo {
	return &userDatabase{
		db: db,
	}
}

func (impl *userDatabase) Insert(ctx context.Context, user types.User) (int64, error) {
	
	if err := impl.db.Table("users").Create(&user).Error; err != nil {
		
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) {
			switch {
			case mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users.uk_email"):
				return 0, ErrDuplicateEntry.WithMessage("邮箱重复")
			case mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users.uk_mobile"):
				return 0, ErrDuplicateEntry.WithMessage("手机号重复")
			default:
				return 0, ErrDuplicateEntry
			}
		}

		return 0, err
	}

	return user.ID, nil 
}


func (impl *userDatabase) FindOne(ctx context.Context, id int64) (*types.User, error) {
	var item types.User

	if err := impl.db.Table("users").Where("id = ?", id).First(&item).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNoFound
		default:
			return nil, err 
		}
	}

	return &item, nil 
}

func (impl *userDatabase) FindOneByEmail(ctx context.Context, email string) (*types.User, error) {
	var item types.User

	err := impl.db.Table("users").Where("email = ?", email).First(&item).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return nil, ErrRecordNoFound
	case nil:
		return &item, nil
	default:
		return nil, err 
	} 
}

func (impl *userDatabase) FindOneByMobile(ctx context.Context, mobile string) (*types.User, error) {
	var resp types.User
	
	if err := impl.db.Table("users").Where("mobile = ?", mobile).First(&resp).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrRecordNoFound
		default:
			return nil, err 
		}
	}

	return &resp, nil 
}

func (impl *userDatabase) Exists(ctx context.Context, id int64) (bool, error) {
	var exists bool

	stmt := "SELECT EXISTS(SELECT true FROM user WHERE id = ?)"

	err := impl.db.Raw(stmt, id).Scan(&exists).Error

	return exists, err
}

func (impl *userDatabase) Delete(ctx context.Context, id int64) error {
	
	if err := impl.db.Transaction(func(tx *gorm.DB) error {
		var exists bool

		stmt := "SELECT EXISTS(SELECT true FROM user WHERE id = ?)"

		if err := tx.Raw(stmt, id).Scan(&exists).Error; err != nil {
			return err
		}
		if !exists {
			return ErrRecordNoFound
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

func (impl *userDatabase) Update(ctx context.Context, user types.User) error {
	return nil
}

func (impl *userDatabase) AllUsers(ctx context.Context) ([]*types.User, error) {
	return nil, nil
}

func (impl *userDatabase) ResetPassword(ctx context.Context, id int64, password string) error {

	return nil 
}


