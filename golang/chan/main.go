package main

import "fmt"

func main() {
	channelTest := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			channelTest <- i
		}
		close(channelTest)
	}()
	if data, ok := <-channelTest; ok {
		fmt.Println(data)
	}
}
