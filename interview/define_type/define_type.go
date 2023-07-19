package main

import (
	"fmt"
	"reflect"
)

// 自定义类型
type FylInt int

func (fyl FylInt) String() string {
	return "fyl"
}

// 别名
type OtherNameInt = int

func main() {
	var a FylInt
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a.String())
	var b OtherNameInt
	fmt.Println(reflect.TypeOf(b))
}
