package utils

import (
	"fmt"
	"math/rand"
)

var numbers    = []rune("0123456789")
var letters    = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var mixLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateCode(size int) string {
	l := make([]rune, 3)
	n := make([]rune, 3)
	for i := range l {
		l[i] = letters[rand.Intn(len(letters))]
	}
	for i := range n {
		n[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(l) + string(n)
}


func RandomString(size int) string {
	m := make([]rune, size)
	for i := range m {
		m[i] = letters[rand.Intn(len(letters))]
	}
	return string(m)
}

func GenerateNum() string {
	// 六位数，num 在 0 ~ 999_999 之间，包括 0 和 9
	num := rand.Intn(1000_000)
	// 不够六位的，加上前导 0
	return fmt.Sprintf("%6d", num)
}


