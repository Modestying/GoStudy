package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8503")
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("hello"))

	var answer [1024]byte
	n, err := conn.Read(answer[:])
	if err != nil {
		fmt.Println("recv failed,err:", err)
		return
	}
	fmt.Println(string(answer[:n]))
	conn.Close()
}
