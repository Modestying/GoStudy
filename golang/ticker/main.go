package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	done := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("xx")
			case <-done:
				fmt.Println("done")
				return
			default:
			}
		}
	}()
	time.Sleep(time.Second * 7)
	done <- struct{}{}
	fmt.Println("end")
	fmt.Println(runtime.NumGoroutine())
}
