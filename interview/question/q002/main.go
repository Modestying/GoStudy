package main

import (
	"fmt"
	"strings"
)

func isUniqueString(data string) bool {
	if strings.Count(data, "") > 3000 {
		return false
	}
	for _, val := range data {
		if val > 127 {
			return false
		}
		if strings.Count(data, string(val)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUniqueString("ASDBASD"))
	fmt.Println(isUniqueString("你好"))
	fmt.Println(isUniqueString("ASDWEOI"))
}
