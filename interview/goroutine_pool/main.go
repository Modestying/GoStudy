package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Task struct {
	Des  string
	Work func()
}
type GPool struct {
	sync.WaitGroup
	taskQueue chan Task
}

func NewGPool(num int) *GPool {
	pool := &GPool{
		taskQueue: make(chan Task),
	}
	pool.Add(num)
	for i := 0; i < num; i++ {
		go pool.Work(i)
	}
	return pool
}

func (pool *GPool) Work(no int) {
	for task := range pool.taskQueue {
		fmt.Println("work id:", no)
		task.Work()
	}
	pool.Done()
}

func (pool *GPool) Stop() {
	close(pool.taskQueue)
	pool.Wait()
}

func (pool *GPool) AddTask(task Task) {
	pool.taskQueue <- task
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
	time.Sleep(time.Second * 1)
	for i := 0; i < 11; i++ {
		i := i
		pool.AddTask(Task{
			Des: strconv.Itoa(i),
			Work: func() {
				fmt.Println(i)
			},
		})
	}
	time.Sleep(time.Second * 4)
}
