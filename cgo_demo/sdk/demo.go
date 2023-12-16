package sdk

/*
#cgo LDFLAGS:-ldemo
#include <stdio.h>
#include <stdlib.h>
#include "demo.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// ProcessInt --> int
func ProcessInt(data int) {
	C.processInt(C.int(data))
}

// ProcessIntPtr --> int *
func ProcessIntPtr(data int) {
	C.processIntPtr((*C.int)(unsafe.Pointer(&data)))
}

// ProcessUnsignedInt --> unsigned int
func ProcessUnsignedInt(data uint) {
	C.processUnsignedInt(C.uint(data))
}

// ProcessUnsignedIntPtr --> unsigned int *
func ProcessUnsignedIntPtr(data uint) {
	C.processUnsignedIntPtr((*C.uint)(unsafe.Pointer(&data)))
}

// ProcessChar --> char
func ProcessChar(data byte) {
	C.processChar(C.char(data))
}

// ProcessCharPtr --> char *
func ProcessCharPtr(data string) {
	str := C.CString(data)
	defer C.free(unsafe.Pointer(str))
	C.processCharPtr(str)
}

// ProcessConstCharPtr --> const char *
func ProcessConstCharPtr(data string) {
	str := C.CString(data)
	defer C.free(unsafe.Pointer(str))
	C.processConstCharPtr(str)
}

// ProcessUnsignedChar --> unsigned char
func ProcessUnsignedChar(data byte) {
	C.processUnsignedChar(C.uchar(data))
}

// ProcessUnsignedCharPtr --> unsigned char *
func ProcessUnsignedCharPtr(data string) {
	str := C.CString(data)
	defer C.free(unsafe.Pointer(str))
	C.processUnsignedCharPtr((*C.uchar)(unsafe.Pointer(str)))
}

// ProcessConstCharPtr --> const unsigned char *
func ProcessConstUnsignedCharPr(data string) {
	str := C.CString(data)
	defer C.free(unsafe.Pointer(str))
	C.processConstUnsignedCharPtr((*C.uchar)(unsafe.Pointer(str)))
}

type Student struct {
	Age  int
	Name [20]byte // [20]byte
}

func (s Student) String() string {
	return fmt.Sprintf("Student: Age = %d, Name = %s-end", s.Age, string(s.Name[:]))
}

// ProcessStruct --> Student
func ProcessStruct(stu Student) {
	fmt.Println(stu)
	cStudent := C.Student{
		Age:  C.int(stu.Age),
		Name: *(*[20]C.char)(unsafe.Pointer(&stu.Name)),
	}
	C.processStruct(cStudent)
}

// processStructPtr --> Student * 这是一个对数组的修改
// 没有发生扩容，切片内存地址不会修改
func ProcessStructPtr(stu []Student, number int) {
	students := make([]C.Student, number)
	C.processStructPtr(&students[0], C.int(number))
	/*
		for i := 0; i < number; i++ {
			age := int(students[i].Age)
			fmt.Printf("Student %d: Age = %d, Name = %s-end\n", i, age, C.GoString(&students[i].Name[0]))
		}
	*/
	for i := 0; i < number; i++ {
		stu[i].Age = int(students[i].Age)
		copy(stu[i].Name[:], []byte(C.GoString(&students[i].Name[0])))
		//fmt.Println(stu[i].String())
	}
}

// ProcessVoidPtr --> void *
func ProcessVoidPtr() {
	a := 10
	ptr := unsafe.Pointer(&a)
	C.processVoidPtr(ptr)
	fmt.Println("result", a) // Output: 42
}

// ProcessVoidPtrPtr --> void **
func ProcessVoidPtrPtr() {
	var ptr unsafe.Pointer
	C.processVoidPtrPtr(&ptr)
	// 通过指针获取具体的值
	value := *(*int)(ptr)
	fmt.Println(value) // Output: 42
	// 通过指针修改具体的值
	*(*int)(ptr) = 89
	C.processVoidPtrPtr(&ptr)
	// 重新获取具体的值
	value = *(*int)(ptr)
	fmt.Println(value) // Output: 42

	// 释放内存，ptr指向的内存空间是在c里面申请的，需要c的free函数释放
	// 可以不调用，程序结束后会自动释放，但不建议这样做，典型的内存泄露问题
	C.free(ptr)

	// 释放后，调用会出错
	// C.processVoidPtrPtr(&ptr)
	// // 重新获取具体的值
	// value = *(*int)(ptr)
	// fmt.Println(value) // Output: 42
}
