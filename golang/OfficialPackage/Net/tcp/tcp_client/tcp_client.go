package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.10.96:8503")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("hello"))

	var answer [1024]byte
	n, err := conn.Read(answer[:])
	if err != nil {
		fmt.Println("rec failed,err:", err)
		return
	}
	fmt.Println(string(answer[:n]))
	conn.Close()
}
