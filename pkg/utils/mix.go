package utils

import "strings"

func Mix(errs []string) string {
	return strings.Join(errs, "、")
}