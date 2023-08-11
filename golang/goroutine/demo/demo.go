package main

import (
	"encoding/binary"
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	x := -1
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, uint16(x))
	fmt.Printf("%b", data)
	fmt.Printf("%x", strconv.FormatFloat(-1, 'x', -1, 64))
	return
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	select {}
}
