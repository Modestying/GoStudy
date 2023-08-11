package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	a int8  //1
	b int64 //8
	c int32 //4
	d int16 //2
}

type T2 struct {
	e int8
	d int64 //8
}

type T3 struct {
	a uint8
	b uint64
}

func main() {
	fmt.Println(unsafe.Sizeof(T{}))
	fmt.Println(unsafe.Sizeof(T2{}))
	fmt.Println(unsafe.Sizeof(T3{}))

}
