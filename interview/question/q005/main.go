package main

import (
	"fmt"
	"strings"
	"unicode"
)

func SubStrWithSpace(str string) (string, bool) {
	if len([]rune(str)) > 1000 {
		return str, false
	}
	for _, val := range str {
		if string(val) != " " && !unicode.IsLetter(val) {
			return str, false
		}
	}
	return strings.Replace(str, " ", "%20", -1), true
}

func main() {
	fmt.Println(SubStrWithSpace("sasd wqw"))
}
