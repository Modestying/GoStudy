package main

import "fmt"

func A() {
	a, b := 1, 2
	defer func() {
		//a = a + 1
		fmt.Println("defer ", a, " ", b)
	}()
	fmt.Println(a, b)
}

func main() {
	A()
}
