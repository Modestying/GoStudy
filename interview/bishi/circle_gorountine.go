package main

import "sync"

func main() {
	wait := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(i int, w *sync.WaitGroup) {
			println(i)
			w.Done()
		}(i, wait)
	}
	wait.Wait()
}
