package data

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"

	"we-backend/pkg/errno"
	"we-backend/internal/types"
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
				return 0, errno.ErrDuplicatedEntry.WithMessage("邮箱重复")
			case mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users.uk_mobile"):
				return 0, errno.ErrDuplicatedEntry.WithMessage("手机号重复")
			default:
				return 0, errno.ErrDuplicatedEntry
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
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, errno.ErrRecordNoFound
	case err != nil :
		return nil, err
	default:
		return &user, nil 
	}
}

func (impl *userRawDatabase) FindOneByEmail(ctx context.Context, email string) (*types.User, error) {
	return nil, nil 
}

func (impl *userRawDatabase) Exists(ctx context.Context, id int64) (bool, error) {
	
	var exists bool

	stmt := "SELECT EXISTS(SELECT true FROM users WHERE id = ?)"

	err := impl.rawDB.QueryRow(stmt, id).Scan(&exists)
	
	return exists, err
}

func (impl *userRawDatabase) Delete(ctx context.Context, id int64) error {
	tx, err := impl.rawDB.Begin()
	if err != nil {
		return err 
	}
	defer tx.Rollback()

	stmt := `DELETE FROM users WHERE id = ?`

	_, err = tx.Exec(stmt, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err 
}

func (impl *userRawDatabase) ResetPassword(ctx context.Context, id int64, password string) error {

	stmt := `UPDATE users SET hashed_password = ? WHERE id = ?`
	
	_, err := impl.rawDB.Exec(stmt, password, id)
	if err != nil {
		return err
	}

	return nil
}

func (impl *userRawDatabase) AllUsers(ctx context.Context) ([]*types.User, error) {
	
	query := `
		SELECT 
			nickname, email, mobile, password, avatar, intro, created_at, updated_at
		FROM 
			users
		ORDER BY 
			nickname DESC
		LIMIT 
			10`

	rows, err := impl.rawDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()


	var users []*types.User

	for rows.Next() {
		var u types.User

		err := rows.Scan(
			&u.Nickname,
			&u.Email,
			&u.Mobile,
			&u.Password,
			&u.Avatar,
			&u.Intro,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}