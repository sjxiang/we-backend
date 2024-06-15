package data

import (
	"context"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"we-backend/internal/biz"
	"we-backend/internal/types"
	"we-backend/pkg/errno"
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
				return 0, errno.ErrDuplicatedEntry.WithMessage("邮箱重复")
			default:
				return 0, errno.ErrDuplicatedEntry
			}
		}

		return 0, err
	}

	return user.ID, nil 
}


func (impl *userDatabase) FindOne(ctx context.Context, id int64) (*types.User, error) {
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

func (impl *userDatabase) FindOneByEmail(ctx context.Context, email string) (*types.User, error) {
	var item types.User

	err := impl.db.Table("users").Where("email = ?", email).First(&item).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return nil, errno.ErrRecordNoFound
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
			return nil, errno.ErrRecordNoFound
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




func (impl *userDatabase) Update(ctx context.Context, user types.User) error {
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

func (impl *userDatabase) AllUsers(ctx context.Context) ([]*types.User, error) {
	return nil, nil
}

func (impl *userDatabase) ResetPassword(ctx context.Context, id int64, password string) error {
	
	return nil 
}


// func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
// 	cacheKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
// 	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cacheKey)

// 	err = func() error {
// 		err1 := cachedResult.Err()
// 		if err1 != nil {
// 			return err1
// 		}
// 		cachedResultByte, err2 := cachedResult.Bytes()
// 		if err2 != nil {
// 			return err2
// 		}
// 		err3 := json.Unmarshal(cachedResultByte, &product)
// 		if err3 != nil {
// 			return err3
// 		}
// 		return nil
// 	}()
// 	if err != nil {
// 		product, err = c.productQuery.GetById(productId)
// 		if err != nil {
// 			return Product{}, err
// 		}
// 		encoded, err := json.Marshal(product)
// 		if err != nil {
// 			return product, nil
// 		}
// 		_ = c.cacheClient.Set(c.productQuery.ctx, cacheKey, encoded, time.Hour)
// 	}
// 	return
// }