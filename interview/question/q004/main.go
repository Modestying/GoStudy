package main

import "strings"

func IsSameStr(str1, str2 string) bool {
	len1 := len([]rune(str1))
	len2 := len([]rune(str2))

	if len1 != len2 {
		return false
	}
	for _, val := range str1 {
		if strings.Count(str1, string(val)) != strings.Count(str2, string(val)) {
			return false
		}
	}
	return true
}
