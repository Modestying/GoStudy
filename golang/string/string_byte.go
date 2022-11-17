package string

import (
	"reflect"
	"unsafe"
)

/*
	1.unsafe.Pointer(&a)方法可以得到变量a的地址
	2.(*reflect.StringHeader)(unsafe.Pointer(&a)) 可以把字符串a转成底层结构的形式
	3.(*[]byte)(unsafe.Pointer(&target)) 可以把target底层结构体转成byte的切片的指针
	4.再通过 *转为指针指向的实际内容
*/

func String2Bytes(str string) []byte {
	//StringHeader是字符串在go的底层结构
	target := (*reflect.StringHeader)(unsafe.Pointer(&str))

	return *(*[]byte)(unsafe.Pointer(&target))
}

func Bytes2String(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
