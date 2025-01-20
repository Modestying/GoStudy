package main

import (
	"flag"
	"runtime"
	"time"
)

var num int
var second int
var sleep int

func main() {
	flag.IntVar(&num, "num", 3, "cpu core nums")
	flag.IntVar(&second, "t", 1000, "ticker")
	flag.IntVar(&sleep, "s", 200, "sleep")

	flag.Parse()
	runtime.GOMAXPROCS(num)

	for i := 0; i < 3; i++ {
		go func() {
			ticker := time.NewTicker(time.Millisecond * time.Duration(second))
			for {
				select {
				case <-ticker.C:
					time.Sleep(time.Millisecond * time.Duration(sleep))
				default:
				}
			}
		}()
	}

	select {}
}
