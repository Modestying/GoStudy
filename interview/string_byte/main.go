package main

import (
	"reflect"
	"unsafe"
)

func ConvertStringToByte(data string) []byte {
	return []byte(data)
}

func ConvertStringToByte2(data string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&data))
	ret := reflect.SliceHeader{Data: str.Data, Len: str.Len, Cap: str.Len}
	return *(*[]byte)(unsafe.Pointer(&ret))
}
func main() {

}
