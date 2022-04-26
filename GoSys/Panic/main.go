package main

import (
	"errors"
	"fmt"
)

//Panic: every goroutine must set recover(eg1).If not , panic(error) will back to goroutine's father thread(eg2)
func main() {
	//eg1
	//TestPanicWithRecover()

	//eg2
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println(r.(error).Error())
	//	}
	//}()
	//PanicWithoutPanic()

	fmt.Println("Main end")
}

func PanicWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r.(error).Error())
		}
	}()

	err := errors.New("PanicWithRecover")
	panic(err)
	return
}

func PanicWithoutPanic() {
	err := errors.New("PanicWithoutPanic")
	panic(err)
	return
}
