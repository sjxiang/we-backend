package types

import (
	"fmt"
	"net/mail"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)



type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`  
}

func (req *LoginRequest) Validate() []string {
	violations := make([]string, 0)
	
	if err := ValidateString(req.Password, 8, 32); err != nil {
		violations = append(violations, err.Error())
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		violations = append(violations, "不是有效的 email")
	}

	if err := validator.New().Var(req.Email, "required,email"); err != nil {
		violations = append(violations, "不是有效的 email")
	}

	allowed := false
	for _, domain := range emailDomainWhitelist {
		if strings.HasSuffix(req.Email, "@"+domain) {
			allowed = true
			break
		}
	}
	if !allowed {
		violations = append(violations, "邮箱地址的域名不在白名单中")
	}

	return violations
}

func ValidateString(value string, minLength int, maxLength int) error {
	n := utf8.RuneCountInString(value)  // 返回字符长度

	if n < minLength || n > maxLength {
		return fmt.Errorf("字符长度必须在 %d ~ %d 之间", minLength, maxLength)
	}
	return nil
}


// 白名单
var emailDomainWhitelist = []string{
	"gmail.com",
	"163.com",
	"126.com",
	"qq.com",
	"outlook.com",
	"hotmail.com",
	"yahoo.com",
	"foxmail.com",
}


type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}


type RegisterRequest struct {
	Email           string   `json:"email"            validate:"required,email"         binding:"required,email"`
	Password        string   `json:"password"         validate:"required,min=8,max=32"  binding:"required,min=8,max=48"`
	PasswordConfirm string   `json:"password_confirm" validate:"eqfield=Password"       binding:"eqfield=Password"`
}
