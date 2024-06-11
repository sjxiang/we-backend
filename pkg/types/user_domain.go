package types

 import (
	"time"
)

type User struct {
	ID        int64      `gorm:"column:id"`     
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`

	Email     string     `gorm:"column:email"`  
	Password  string     `gorm:"column:password"`     
	Nickname  string     `gorm:"column:nickname"`
	Mobile    string     `gorm:"column:mobile"`   
	Intro     string     `gorm:"column:intro"`  
	Avatar    string     `gorm:"column:avatar"`
	// YYYY-MM-DD
	Birthday  WrapTime   `gorm:"column:birthday"`
}

func (u User) TableName() string {
	return "users"
}

// TodayIsBirthday 判定今天是不是我的生日
func (u User) TodayIsBirthday() bool {
	now := time.Now()
	return now.Month() == u.Birthday.Month() && now.Day() == u.Birthday.Day()
}

// type UserM struct {
// 	// 自增主键
// 	ID        int64             `gorm:"primaryKey,autoIncrement"`

// 	// 设置为唯一索引
// 	Email     string            `gorm:"uniqueIndex;type:varchar(128);not null"`
// 	Password  string            `gorm:"type:varchar(64);not null"`


// 	// 代表这是一个可以为 NULL 的列
// 	Nickname  string            `gorm:"type:varchar(128)"`	
// 	Mobile    *string           `gorm:"type:varchar(11)"`
// 	Intro     string            `gorm:"type:varchar(4096)"`

// 	// 代表这是一个可以为 NULL 的列
// 	Avatar    sql.NullString    `gorm:"type:varchar(4096)"`

// 	// YYYY-MM-DD，UTC 0 的毫秒数
// 	Birthday  int64             `gorm:"type:varchar(4096)"`

// 	Gender    string            `gorm:"column:gender;type:varchar(6);default:male;comment 'female表示女，male表示男'"`
// 	Role      int               `gorm:"column:role;type:int;default:1;comment '1表示普通用户，2表示管理员'"`

// 	CreatedAt int64             `gorm:"column:created_at;not null;comment '创建时间，毫秒数'"`
// 	UpdatedAt int64				`gorm:"column:updated_at;not null"`

// 	DeletedAt gorm.DeletedAt    `gorm:"index"`
// }

// func (u *UserM) TableName() string {
// 	return consts.TABLE_NAME_USER
// }

