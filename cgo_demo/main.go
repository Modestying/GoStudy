package main

import (
	"cgo_example/sdk"
)

func main() {
	// a := 1
	// sdk.ProcessInt(a)
	// sdk.ProcessIntPtr(a)

	// sdk.ProcessUnsignedInt(uint(a))
	// sdk.ProcessUnsignedIntPtr(uint(a))

	// sdk.ProcessChar('a')
	// sdk.ProcessCharPtr("ProcessCharPtr")
	// sdk.ProcessConstCharPtr("ProcessConstCharPtr")

	// sdk.ProcessUnsignedChar('a')
	// sdk.ProcessUnsignedCharPtr("ProcessUnsignedCharPtr")
	// sdk.ProcessConstUnsignedCharPr("ProcessConstUnsignedCharPr")
	// stu := sdk.Student{
	// 	Age: 131,
	// }
	// copy(stu.Name[:], []byte("0123456789abcdefghijk"))
	// sdk.ProcessStruct(stu)

	// students := make([]sdk.Student, 4)
	// sdk.ProcessStructPtr(students, 4)
	// for _, v := range students {
	// 	fmt.Println(v.String())
	// }

	//sdk.ProcessVoidPtr()
	sdk.ProcessVoidPtrPtr()
}
