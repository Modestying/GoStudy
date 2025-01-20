package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Person struct {
	name string
	age  int
}

// 全局变量（简单处理）
var p atomic.Value

func update(name string, age int) {
	lp := &Person{}
	// 更新第一个字段
	lp.name = name
	// 加点随机性
	time.Sleep(time.Millisecond * 200)
	// 更新第二个字段
	lp.age = age
	// 原子设置到全局变量
	p.Store(lp)
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	// 10 个协程并发更新
	for i := 0; i < 10; i++ {
		name, age := fmt.Sprintf("nobody:%v", i), i
		go func() {
			defer wg.Done()
			update(name, age)
		}()
	}
	wg.Wait()
	// 结果是啥？你能猜到吗？
	_p := p.Load().(*Person)
	fmt.Printf("p.name=%s\np.age=%v\n", _p.name, _p.age)
}
