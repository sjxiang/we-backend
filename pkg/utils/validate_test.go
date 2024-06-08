package utils

import (
	"strings"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	minSize, digit, special, letter := ValidatePassword("hasicoghwif*4YY")
	if !minSize || !digit || !special || !letter {
		t.Log("无效密码")
		msg := "密码："
		var errs []string
		if !minSize {
			errs = append(errs, "最少 8 个字符")
		}
		if !digit {
			errs = append(errs, "至少要有数字")
		}
		if !special {
			errs = append(errs, "至少要有特殊字符")
		}
		if !letter {
			errs = append(errs, "至少要有字母")
		}
		
		t.Log(msg + strings.Join(errs, "，"))
	}
}
