package main

import (
	"fmt"
	"runtime"
	"time"
)

var x string

func do2phases(phase1, done chan bool) {
	time.Sleep(time.Second * 3) // 第 1 段
	select {
	case phase1 <- true:
		fmt.Println("phase1 success")
	default:
		fmt.Println("xxxx")
		return
	}
	time.Sleep(time.Second) // 第 2 段
	done <- true
}

func timeoutFirstPhase() error {
	phase1 := make(chan bool)
	done := make(chan bool)
	go do2phases(phase1, done)
	select {
	case <-phase1:
		<-done
		fmt.Println("done")
		return nil
	case <-time.After(time.Second * 1):
		fmt.Println("timeout phase")
		return fmt.Errorf("timeout")
	}
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		timeoutFirstPhase()
	}
	time.Sleep(time.Second * 10)
	fmt.Println(runtime.NumGoroutine())
	if time.Since(start) > time.Second*1 {
		fmt.Println("xx")
	}
}
