package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	x := make([]int, 2)
	x = append(x, 1)
	fmt.Println(x)
	return
	runtime.GOMAXPROCS(2)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("xx")
	}
}
