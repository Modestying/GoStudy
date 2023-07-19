package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var ctx, cancel = context.WithCancel(context.Background())

type item struct {
	val        string
	recordTime time.Time
}
type SafeMap struct {
	m   map[string]*item
	t   *time.Ticker
	ctx context.Context
	sync.RWMutex
}

func NewSafeMap(ctx context.Context) *SafeMap {
	instance := &SafeMap{
		m:   make(map[string]*item, 10),
		t:   time.NewTicker(time.Second * 5),
		ctx: ctx,
	}
	go instance.cleanUp()
	return instance
}

func (s *SafeMap) Close() {
	s.t.Stop()
	cancel()
	time.Sleep(time.Second * 1)
}
func (s *SafeMap) cleanUp() {
	for {
		select {
		case <-s.t.C:
			s.Lock()
			for k, v := range s.m {
				if time.Since(v.recordTime) > time.Second*5 {
					delete(s.m, k)
				}
			}
			s.Unlock()
			s.t.Reset(time.Second * 5)
		case <-s.ctx.Done():
			for k, _ := range s.m {
				delete(s.m, k)
			}
			fmt.Println("cleanUp done")
			return
		}
	}
}
func (s *SafeMap) Set(key, val string) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = &item{
		val:        val,
		recordTime: time.Now(),
	}
}

func (s *SafeMap) Get(key string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	data, ok := s.m[key]
	if ok {
		return data.val, true
	} else {
		return "", false
	}
}

func main() {
	s := NewSafeMap(ctx)
	s.Set("a", "1")
	s.Set("b", "2")
	s.Set("c", "3")

	fmt.Println(s.Get("a"))
	fmt.Println(s.Get("b"))
	fmt.Println(s.Get("c"))

	s.Close()

	fmt.Println(s.Get("a"))
}
