package types

import (
	"fmt"
	"net/mail"
	"strings"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`  
}

func (req *LoginRequest) Validate() []string {
	errs := make([]string, 0)
	
	if err := ValidateString(req.Password, 8, 32); err != nil {
		errs = append(errs, err.Error())
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		errs = append(errs, "不是有效的 email")
	}

	if err := validator.New().Var(req.Email, "required,email"); err != nil {
		errs = append(errs, "不是有效的 email")
	}

	allowed := false
	for _, domain := range emailDomainWhitelist {
		if strings.HasSuffix(req.Email, "@"+domain) {
			allowed = true
			break
		}
	}
	if !allowed {
		errs = append(errs, "邮箱地址的域名不在白名单中")
	}

	return errs
}

func (req *LoginRequest) ExportEmailInString() string {
	return req.Email
}

func (req *LoginRequest) ExportPasswordInString() string {
	return req.Password
}


func ValidateString(value string, minLength int, maxLength int) error {
	// n := len(value)  // 返回字节长度
	
	n := utf8.RuneCountInString(value)  // 返回字符长度

	if n < minLength || n > maxLength {
		return fmt.Errorf("字符长度必须在 %d-%d 之间", minLength, maxLength)
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

}

func (x *LoginResponse) ExportForFeedback() *LoginResponse {
	return x
}
