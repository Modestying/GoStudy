package main

import (
	"errors"
	"fmt"
)

// Panic: every goroutine must set recover(eg1).If not , panic(error) will back to goroutine's father thread(eg2)

// PanicWithRecover 在函数内部实现recover,内部处理错误
func PanicWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Msg:", r.(error).Error())
		}
	}()

	err := errors.New("PanicWithRecover")
	panic(err)
}

// PanicWithoutRecover 函数无recover，会跳到上层函数寻找recover
func PanicWithoutRecover() {

	fmt.Println("PanicWithoutPanic....  ")
	err := errors.New("PanicWithoutPanic")
	panic(err)
}

func proc() {
	panic("panic proc")
}
func PanicContinue() {
	for i := 0; i < 10; i++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic Msg:", r)
				}
			}()
			proc()
		}()
	}
}

// go run panic.go > output.txt 2>&1
func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Panic Msg:", r.(error).Error())
	// 	}
	// }()
	fmt.Println("xx")
	PanicContinue()
	//PanicWithoutRecover()
}
