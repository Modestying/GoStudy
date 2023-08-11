package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	runtime.GOMAXPROCS(1)
	file, _ := os.Create("trace.out")
	defer file.Close()

	trace.Start(file)
	fmt.Println("hello,gmp")
	trace.Stop()
}
