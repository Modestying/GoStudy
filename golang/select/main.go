package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 1)
		close(ch1)
	}()
	select {
	case _, ok := <-ch1:
		if ok {
			fmt.Println("ss")
		} else {
			fmt.Println("SSSS")
		}
	case <-ch2:
		fmt.Println("xx")
	case <-time.After(time.Second * 6):
		fmt.Println("timeout")
		return
	}
}
