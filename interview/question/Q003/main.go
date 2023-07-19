package main

import (
	"strings"
)

func ReverseString(data string) (string, bool) {
	if strings.Count(data, "") > 5000 {
		return "", false
	}
	chars := []rune(data)
	length := len(chars)
	for i := 0; i < length/2; i++ {
		chars[i], chars[length-i-1] = chars[length-i-1], chars[i]
	}
	return string(chars), true
}
