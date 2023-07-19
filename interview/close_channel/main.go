package main

import (
	"fmt"
	"sync"
)

// 安全的关闭channel

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("发生了 panic:", err)
		}
	}()
	channelA := make(chan int)
	close(channelA)
	channelB := make(chan int)
	instance := sync.Once{}
	instance.Do(func() {
		close(channelB)
	})
	close(channelA)
}
