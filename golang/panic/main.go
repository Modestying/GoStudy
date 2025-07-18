package main

import (
	"runtime/debug"
	"time"
)

func main() {
	debug.SetTraceback("crash")

	time.Sleep(time.Second)
	panic("xx")
}
