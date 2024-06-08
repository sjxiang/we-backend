package utils

import (
	"unicode"
	"regexp"
)


// 校验密码组成，组合（数字 + 字母 + 特殊字符）
func ValidatePassword(password string) (minSize, digit, special, letter bool) {
	for _, c := range password {
		switch {
		// 数字
		case unicode.IsNumber(c):
			digit = true
		// 字母（大小写不敏感）
		case unicode.IsUpper(c) || unicode.IsLower(c):
			letter = true
		// 特殊字符
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}
	minSize = len(password) >= 8
	return
}


var (
	nameRegexp  = regexp.MustCompile(`^[a-z][a-z0-9-]{0,39}$`)
)

// 小写字母开头，后面可以跟小写字母、数字或破折号-，最多允许 40 个字符的长度。
func ValidateNickname(name string) bool {
	return nameRegexp.MatchString(name)
}