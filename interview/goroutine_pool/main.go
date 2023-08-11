package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type GPool struct {
	controler chan struct{}
}

func NewGPool(num int) *GPool {
	pool := &GPool{
		controler: make(chan struct{}, num),
	}
	for i := 0; i < num; i++ {
		pool.controler <- struct{}{}
	}
	return pool
}

func (pool *GPool) Do(f func(any), data any) {
	<-pool.controler
	f(data)
	pool.controler <- struct{}{}
}
func main1() {
	ch := make(chan struct{}, 5)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		ch <- struct{}{}
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			// 阻塞等待一个数据传入
			<-ch
			time.Sleep(time.Second)
			fmt.Println("finish job ", i)
			defer func() {
				ch <- struct{}{}
				wg.Done()
			}()
		}()
	}

	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
}

func main() {
	pool := NewGPool(5)
	for i := 0; i < 11; i++ {
		i := i
		go pool.Do(func(data any) {
			time.Sleep(time.Second)
			fmt.Println("job index:", data)
		}, i)
	}
	select {}
}
