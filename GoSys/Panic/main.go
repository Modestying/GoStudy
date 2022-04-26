package main

import (
	"errors"
	"fmt"
)

func main() {
	TestPanic()
	fmt.Println("a")
}

func TestPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ss")
		}
	}()

	err := errors.New("dd")
	panic(err)
	return
}
