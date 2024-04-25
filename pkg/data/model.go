package data

import (
	"database/sql"
	"we-backend/pkg/consts"

	"gorm.io/gorm"
)


type UserM struct {
	// 自增主键
	ID        int64             `gorm:"primaryKey,autoIncrement"`

	// 设置为唯一索引
	Email     string            `gorm:"uniqueIndex;type:varchar(128);not null"`
	Password  string            `gorm:"type:varchar(64);not null"`


	// 代表这是一个可以为 NULL 的列
	Nickname  string            `gorm:"type:varchar(128)"`	
	Mobile    *string           `gorm:"type:varchar(11)"`
	Intro     string            `gorm:"type:varchar(4096)"`

	// 代表这是一个可以为 NULL 的列
	Avatar    sql.NullString    `gorm:"type:varchar(4096)"`

	// YYYY-MM-DD，UTC 0 的毫秒数
	Birthday  int64             `gorm:"type:varchar(4096)"`

	Gender    string            `gorm:"column:gender;type:varchar(6);default:male;comment 'female表示女，male表示男'"`
	Role      int               `gorm:"column:role;type:int;default:1;comment '1表示普通用户，2表示管理员'"`

	CreatedAt int64             `gorm:"column:created_at;not null;comment '创建时间，毫秒数'"`
	UpdatedAt int64				`gorm:"column:updated_at;not null"`

	DeletedAt gorm.DeletedAt    `gorm:"index"`
}

func (u *UserM) TableName() string {
	return consts.TABLE_NAME_USER
}

