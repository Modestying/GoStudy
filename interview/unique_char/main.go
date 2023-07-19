package main

import (
	"fmt"
	"strings"
)

func UniqueChar(data string) bool {
	for _, v := range data {
		if strings.Count(data, string(v)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(UniqueChar("abcdea"))
	fmt.Println(UniqueChar("你好啊你"))
}
