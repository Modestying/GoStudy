package main

import (
	"time"
)

func main() {
	t := time.NewTicker(time.Second * 1)
	go func(t *time.Ticker) {
		time.Sleep(time.Second * 5)
		t.Stop()
	}(t)
	// for _ = range t.C {
	// 	fmt.Println("tick")
	// }
	for {
		if _, ok := <-t.C; !ok {
			break
		} else {
			println("tick")
		}
	}
}
