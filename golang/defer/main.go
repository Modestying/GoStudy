package main

import (
	"fmt"
	"time"
)

func A() { //
	a, b := 1, 2
	defer func() {
		a = a + 1
		fmt.Println("defer ", a, " ", b)
	}()
	fmt.Println(a, b)
}

func DeferGO() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 1)
}

// defer在执行ret时会把值复制给ret的变量
func DeferA() int { // 返回10
	a := 10
	defer func() {
		a = 20
	}()
	return a
}

func DeferB() (a int) { // 返回30
	a = 10
	defer func() {
		a = 30
	}()
	return a
}

func DeferD() (a int) { // 返回10
	a = 10
	defer func(a int) {
		a = 30
	}(a)
	return a
}

func Defer3() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) // 传进来的i是地址
		}()
	}
}

func Defer1() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // 传进来的值会进行复制
	}
}

func Defer2() {
	for i := 0; i < 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func Aa() {
	fmt.Println("aa")
}
func main() {
	// Defer1()
	// Defer2()
	// Defer3()
	// fmt.Println(DeferA())
	// fmt.Println(DeferB())
	// fmt.Println(DeferD())
	defer func() {
		recover()
		Aa()
	}()
	panic("xx")
}

func sum(a int) int {
	return a + 10
}
func SumWithDefer() {
	defer func() {
		sum(11)
	}()
}

func SumWithoutDefer() {
	sum(11)
}
