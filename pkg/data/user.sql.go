package data

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"we-backend/pkg/biz"
	"we-backend/pkg/consts"
	"we-backend/pkg/types"
	"we-backend/pkg/x"
)


type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) biz.UserRepo {
	return &userDatabase{
		DB: db,
	}
}

func (impl *userDatabase) Insert(ctx context.Context, arg types.CreateUserParams) error {

	now := time.Now().UnixMilli()

	var item UserM

	item.Email     = arg.Email
	item.Password  = arg.Password
	item.CreatedAt = now
	item.UpdatedAt = now

	if err := impl.DB.Table(consts.TABLE_NAME_USER).Create(&item).Error; err != nil {
		
		// ERROR 1062 (23000): Duplicate entry 'xxx' for key 'xxx'
		// 'user.uk_email'
		// 'user.uk_mobile'
		// 'user.uk_user_id'

		var mysqlError *mysql.MySQLError

		if errors.As(err, &mysqlError) {
			if mysqlError.Number == 1062 {
				switch {
				case strings.Contains(mysqlError.Message, "user.uk_email"):
					return x.ErrDuplicateEmail
				case strings.Contains(mysqlError.Message, "user.uk_mobile"):
					return x.ErrDuplicateMobile
				case strings.Contains(mysqlError.Message, "user.uk_user_id"):
					return x.ErrDuplicateUserID
				}
			}	
		}

		return err
	}
	
	return nil
}


const (
	createUser = `
INSERT INTO user 
	(email, password, created_at, updated_at) 
VALUES 
	(?, ?, ?, ?)
	`
)

func (impl *userDatabase) CreateUser(ctx context.Context, arg types.CreateUserParams) error {
	query := "INSERT INTO `user` (`email`,`password`,`created_at`,`updated_at`) VALUES (?, ?, ?, ?)"

	now := time.Now().UnixMilli()

	if err := impl.DB.Exec(query, arg.Email, arg.Password, now, now).Error; err != nil {
	
		const uniqueViolation uint16 = 1062  

		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case uniqueViolation: 
				return x.ErrDuplicateEmail
			}
		}
		
		return err
	}

	return nil
}


func (impl *userDatabase) Exists(email string) (bool, error) {
	var exists bool
		
	query := `
	SELECT EXISTS(
		SELECT true FROM account WHERE email = ?
	)
	`
	
	if err := impl.DB.Raw(query, email).Scan(&exists).Error; err != nil {
		return false, err
	}

	if !exists {
		return false, nil 
	}

	return true, nil 
}


func (impl *userDatabase) FindUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var row types.User

	if err := impl.DB.Table(consts.TABLE_NAME_USER).Where("email = ?").First(&row).Error; err != nil {
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, x.ErrNoRecord
		} 		
		
		return nil, err
	}

	return &types.User{
		ID:       row.ID,
		Email:    row.Email,
		Password: row.Password,
	}, nil 
}	


func (impl *userDatabase) FindUserByMobile(ctx context.Context, mobile string) (types.User, error) {
	var item types.User

	if err := impl.DB.Table(consts.TABLE_NAME_USER).Where("mobile = ?").First(&item).Error; err != nil {
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return types.User{}, x.ErrNoRecord
		} 		
		
		return types.User{}, err
	}

	return types.User{
		ID:       item.ID,
		Email:    item.Email,
		Password: item.Password,
	}, nil 
}	



// FindUserByUserID(ctx context.Context, userID uint) (user domain.User, err error)
// FindUserByEmail(ctx context.Context, email string) (user domain.User, err error)
// FindUserByUserName(ctx context.Context, userName string) (user domain.User, err error)
// FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (user domain.User, err error)
// FindUserByUserNameEmailOrPhoneNotID(ctx context.Context, user domain.User) (domain.User, error)
// SaveUser(ctx context.Context, user domain.User) (userID uint, err error)
// UpdateVerified(ctx context.Context, userID uint) error
// UpdateUser(ctx context.Context, user domain.User) (err error)
// UpdateBlockStatus(ctx context.Context, userID uint, blockStatus bool) error


