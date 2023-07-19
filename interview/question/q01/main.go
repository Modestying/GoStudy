package main

import (
	"fmt"
)

func Print() {
	chanNum := make(chan struct{}, 1)
	chanChar := make(chan struct{}, 1)

	num := 1
	char := 'A'
	chanNum <- struct{}{}

	for {
		select {
		case <-chanNum:
			fmt.Printf("%d%d", num, num+1)
			num += 2
			if num >= 28 {
				return
			}
			chanChar <- struct{}{}
		case <-chanChar:
			fmt.Printf("%c%c", char, char+1)
			char += 2
			chanNum <- struct{}{}
		}
	}
}

func main() {
	Print()
}
