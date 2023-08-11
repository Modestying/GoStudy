package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	for i := 0; i < 258; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
