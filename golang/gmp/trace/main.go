package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	sync.Once
	runtime.GOMAXPROCS(1)
	file, _ := os.Create("trace.out")
	defer file.Close()

	trace.Start(file)
	fmt.Println("hello,gmp")
	trace.Stop()
}
