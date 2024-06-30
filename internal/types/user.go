package types

import (
	"time"
)

type User struct {
	ID        int64      `gorm:"column:id"          json:"id"`         
	CreatedAt time.Time  `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"  json:"-"`

	Email     string     `gorm:"column:email"`  
	Password  string     `gorm:"column:password"    json:"-"`     
	Nickname  string     `gorm:"column:nickname"`
	Mobile    string     `gorm:"column:mobile"`   
	Intro     string     `gorm:"column:intro"`  
	Avatar    string     `gorm:"column:avatar"`
	Birthday  int64      `gorm:"column:birthday"` // YYYY-MM-DD
}

func (u User) TableName() string {
	return "users"
}

// TodayIsBirthday 判定今天是不是我的生日
func (u User) TodayIsBirthday() bool {
	birth := time.Unix(u.Birthday, 0)
	return time.Now().Month() == birth.Month() && time.Now().Day() == birth.Day()
}


type UserM struct {

}

func (u UserM) TableName() string {
	return "users"
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


type MeInput struct {
	UserID int64 
}

type MeResponse struct {
	User User `json:"user"`
}

type EditRequest struct {
	Nickname string `json:"nickname" validate:"required,min=8,max=30"`
	Birthday string `json:"birthday" validate:"required,len=10"`  // 1997-09-12
	Intro    string `json:"intro"    validate:"required,min=1,max=1000"`
	Avatar   string `json:"avatar"   validate:"required"`
}


type EditInput struct {
	UserID   int64
	Nickname string 
	Birthday int64 
	Intro    string 
	Avatar   string 
}


// tips 小贴士 修改密码、邮箱、手机，都要二次验证
type ResetPasswordRequest struct {
	Password        string   `json:"password"         validate:"required,min=8,max=32"`
	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"`
}


type AllResponse struct {
	Users []*User `json:"users"`
}