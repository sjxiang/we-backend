package types

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
	"we-backend/pkg/we"
)


var (
	UsernameMinLength = 2
	PasswordMinLength = 6
)


var (
	emailRegexp = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)


type RegisterInput struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// 预处理
func (in *RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)

	in.Username = strings.TrimSpace(in.Username)
}

// 校验
func (in RegisterInput) Validate() error {
	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf("%w: username not long enough, (%d) characters at least", we.ErrInvalidParameter, UsernameMinLength)
	}

	if !emailRegexp.MatchString(in.Email) {
		return fmt.Errorf("%w: email not valid", we.ErrInvalidParameter)
	}

	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: password not long enough, (%d) characters at least", we.ErrInvalidParameter, PasswordMinLength)
	}

	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: confirm password must match the password", we.ErrInvalidParameter)
	}

	return nil
}


type RegisterResponse struct {
	User User `json:"user"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (in *LoginInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
}

func (in LoginInput) Validate() error {
	if !emailRegexp.MatchString(in.Email) {
		return fmt.Errorf("%w: email not valid", we.ErrInvalidParameter)
	}

	if len(in.Password) < 1 {
		return fmt.Errorf("%w: password required", we.ErrInvalidParameter)
	}

	return nil
}



// 校验，密码成分（阿拉伯数字 + 英文字母，大小写不敏感 + 特殊字符）密码长度（大于 8）
func ValidatePassword(password string) (minSize, digit, special, letter bool) {
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			digit = true
		case unicode.IsUpper(c) || unicode.IsLower(c):
			letter = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}
	
	minSize = utf8.RuneCountInString(password) >= 8  
	return
}


type SentOtpInput struct {
	Biz         string
	PhoneNumber string
}


type VerifyOtpInput struct {
	Biz         string
	PhoneNumber string
	InputCode   string
}

type SentOtpRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`  // 实际是邮箱
}

type VerifyOtpRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`  // 实际是邮箱
	InputCode   string `json:"input_code" validate:"required"`
}
