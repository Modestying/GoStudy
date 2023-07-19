package main

import (
	"fmt"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type item struct {
	Val interface{}
	Ch  chan struct{}
}
type FMap struct {
	cache map[string]*item
	sync.RWMutex
}

func NewFMap() *FMap {
	instance := &FMap{
		cache: make(map[string]*item, 10),
	}
	return instance
}

func (f *FMap) Out(key string, val interface{}) {
	f.Lock()
	defer f.Unlock()
	if v, ok := f.cache[key]; ok {
		v.Val = val
		if v.Ch != nil {
			close(v.Ch)
			v.Ch = nil
		}
		return
	}
	f.cache[key] = &item{
		Val: val,
	}
}

func (f *FMap) Rd(key string, timeout time.Duration) interface{} {
	f.Lock()
	if v, ok := f.cache[key]; ok {
		f.Unlock()
		return v.Val
	} else {
		it := &item{
			Ch: make(chan struct{}, 1),
		}
		f.cache[key] = it
		f.Unlock()
		select {
		case <-it.Ch:
			return it.Val
		case <-time.After(timeout):
			return nil
		}
	}
}
func main() {
	f := NewFMap()
	go func() {
		fmt.Println("11", f.Rd("11", time.Second*5))
	}()
	go func() {
		fmt.Println("22", f.Rd("11", time.Second*5))
	}()
	f.Out("11", "xx")
	fmt.Println(f.Rd("11", time.Second*5))

	time.Sleep(time.Second * 6)
}
