package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func ElegantClose(wait *sync.WaitGroup) {
	c := make(chan os.Signal)
	signal.Notify(c)
	for data := range c {
		fmt.Println("Receive Signal :", data.String())
		switch data.String() {
		case "interrupt":
			close(c)
			break
		default:
			fmt.Println("default")
			break
		}
	}

	if _, ok := <-c; !ok {
		fmt.Println("Success Close Signal Channel")
	}

	wait.Done()
}

// 监听全部信号
func main() {
	//合建chan
	//监听所有信号
	var wait sync.WaitGroup
	wait.Add(1)
	go ElegantClose(&wait)
	wait.Wait()
	fmt.Println("Service Exit")
}
