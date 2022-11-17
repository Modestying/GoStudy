package string

import (
	"strings"
)

func Equal(a, b string) bool {
	return a == b
}

func Compare(a, b string) bool {
	if strings.Compare(a, b) == 0 {
		return true
	} else {
		return false
	}
}

func EqualFold(a, b string) bool {
	if strings.EqualFold(a, b) {
		return true
	} else {
		return false
	}
}
