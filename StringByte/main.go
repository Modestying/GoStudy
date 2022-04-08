package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	1.unsafe.Pointer(&a)方法可以得到变量a的地址
	2.(*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式
	3.(*[]byte)(unsafe.Pointer(&ssh)) 可以把ssh底层结构体转成byte的切片的指针
	4.再通过 *转为指针指向的实际内容
*/

func main() {
	a := "aaa"
	a = "bbb"
	//StringHeader是字符串在go的底层结构

	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))

	b := *(*[]byte)(unsafe.Pointer(&ssh))
	fmt.Printf("%v", b)
	fmt.Println(a)

	//eg:
	str := "asong"
	by := []byte(str)
	fmt.Println(by)
	str1 := string(by)
	fmt.Println(str1)
}
