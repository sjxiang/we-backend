package utils

import "strings"

func Mix(violations []string) string {
	return strings.Join(violations, "、")
}