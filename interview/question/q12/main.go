package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {

		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println("panic error:", err)
						}
					}()
					proc()
				}()
			}
		}
	}()
	select {}
}

func proc() {
	panic("ok")
}
