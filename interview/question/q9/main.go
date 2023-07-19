package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	chTrans := make(chan int, 5)
	wait := &sync.WaitGroup{}
	wait.Add(1)
	go func(ch chan int) {
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(5)
		}
		close(ch)
	}(chTrans)

	go func(ch chan int, w *sync.WaitGroup) {
		for data := range ch {
			fmt.Println(data)
		}
		w.Done()
	}(chTrans, wait)
	wait.Wait()
}
