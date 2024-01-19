package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	wait := &sync.WaitGroup{}
	wait.Add(1)
	c := make(chan os.Signal, 1)
	defer func() {
		fmt.Println("defer close")
	}()
	signal.Notify(c)
	go func() {
		for data := range c {
			fmt.Println("Receive Signal :", data.String())
			switch data.String() {
			case "interrupt":
				close(c)
				wait.Done()
			default:
				fmt.Println("default")
			}
		}
	}()
	fmt.Println("Start listen interrupt")
	wait.Wait()
}
