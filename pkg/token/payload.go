package token

import (
	"errors"
	"time"
)

var (
	ErrInvalidToken = errors.New("token is invalid")   // 无效
	ErrExpiredToken = errors.New("token has expired")  // 过期
)

// 载体
type Payload struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`   // 签发 YY-MM-DD 
	ExpiredAt time.Time `json:"expired_at"`  // 截至
}

func NewPayload(id int64, email string, duration time.Duration) *Payload {
	payload := &Payload{
		ID:        id,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration), 
	}
	return payload
}


func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}


/*


type Claims interface {
	Valid() error
}

v4 需要实现 jwt.Claims 接口

 */
