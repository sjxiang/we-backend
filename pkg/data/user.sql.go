package data

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"we-backend/pkg/types"

	"github.com/go-sql-driver/mysql"
)


type userRawDatabase struct {
	rawDB *sql.DB
}

func (impl *userRawDatabase) Insert(ctx context.Context, user types.User) (int64, error) {
	
	stmt := `
		INSERT INTO 
			users (nickname, email, mobile, password, avatar, intro, created_at, updated_at)
    	VALUES
			(?, ?, ?, ?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())`

	result, err := impl.rawDB.Exec(
		stmt, 
		user.Nickname, 
		user.Email, 
		user.Mobile,
		user.Password, 
		user.Avatar,
		user.Intro)

	if err != nil {
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

	newID, err := result.LastInsertId()
	if err != nil {
		return 0, err 
	}
	
	return newID, nil 
}

func (impl *userRawDatabase) FindOne(ctx context.Context, id int64) (*types.User, error) {

	query := `
		SELECT 
			nickname, email, mobile, password, avatar, intro, created_at, updated_at
		FROM 
			users
		WHERE 
			id = ?
		LIMIT 
			1`
			
	var user types.User

	row := impl.rawDB.QueryRow(query, id)
	
	err := row.Scan(
		&user.ID,
		&user.Nickname, 
		&user.Email, 
		&user.Mobile,
		&user.Password, 
		&user.Avatar,
		&user.Intro,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNoFound
		} else {
			return nil, err
		}
	}

	return &user, nil
}

func (impl *userRawDatabase) FindOneByEmail(ctx context.Context, email string) (*types.User, error) {
	return nil, nil 
}

func (impl *userRawDatabase) Exists(id int64) (bool, error) {
	
	var exists bool

	stmt := "SELECT EXISTS(SELECT true FROM user WHERE id = ?)"

	err := impl.rawDB.QueryRow(stmt, id).Scan(&exists)
	
	return exists, err
}