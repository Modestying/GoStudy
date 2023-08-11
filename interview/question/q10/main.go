package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{}, tag string)                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration, tag string) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type item struct {
	Val   interface{}
	Ch    chan struct{}
	IsSet bool
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

func (f *FMap) Out(key string, val interface{}, tag string) {
	f.Lock()
	defer f.Unlock()
	if v, ok := f.cache[key]; ok {
		v.Val = val
		v.IsSet = true
		if v.Ch != nil {
			log.Println("close channel ", tag)
			close(v.Ch)
			v.Ch = nil
		}
		return
	}
	f.cache[key] = &item{
		Val:   val,
		IsSet: true,
	}
}

func (f *FMap) Rd(key string, timeout time.Duration, tag string) interface{} {
	instance := &item{
		Ch:    make(chan struct{}, 1),
		IsSet: false,
	}
	f.Lock()
	if it, ok := f.cache[key]; ok {
		f.Unlock()
		instance = it
		if it.IsSet {
			return it.Val
		}
		log.Println("get channel ", tag)
	} else {
		f.cache[key] = instance
		log.Println("set default data:", key, " ", tag)
		f.Unlock()
	}
	log.Println("select wait:....", tag)
	select {
	case <-instance.Ch:
		log.Println("get:", instance.Val, " tag:", tag)
		return instance.Val
	case <-time.After(timeout):
		return nil
	}
}
func main1() {
	f := NewFMap()
	log.Println("start...")
	go func() {
		time.Sleep(time.Second * 3)
		log.Println("first get ", f.Rd("11", time.Second*1, "first"))
	}()
	go func() {
		time.Sleep(time.Second * 4)
		go f.Rd("11", time.Second*3, "second")
		go f.Rd("11", time.Second*3, "third")
	}()
	go func() {
		time.Sleep(time.Second * 5)
		f.Out("11", "xx", "set")
	}()
	// go func() {
	// 	time.Sleep(time.Second * 7)
	// 	f.Out("11", "xx")
	// }()

	time.Sleep(time.Second * 13)
	fmt.Println("xx")
}

func d(ch chan int) {
	ch <- 1
}
func main() {
	ch := make(chan int, 1)
	d(ch)
	<-ch
	runtime.GC()
}
