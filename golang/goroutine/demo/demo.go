package main

import (
	"fmt"
	"sync"
)

func main() {
	wait := &sync.WaitGroup{}
	wait.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			wait.Done()
		}()
	}
	wait.Wait()

}
