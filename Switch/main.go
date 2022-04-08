package main

import (
	"fmt"
)

/*
	1.Golang的case语句中不需要加break
	2.case匹配到的函数结束后，可以使用fallthrough关键字强制执行下一个case中语句
*/
func main() {
	data := 1
	switch data {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough
	default:
		fmt.Println("default")
	}
}
