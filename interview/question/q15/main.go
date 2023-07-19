package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	w := &sync.WaitGroup{}

	go func() {
		w.Add(1)

		w.Wait()
		fmt.Println("xxx")
	}()

	go func() {
		time.Sleep(time.Second * 2)
		w.Done()
		w.Add(1)
	}()

	time.Sleep(time.Second * 5)
}
