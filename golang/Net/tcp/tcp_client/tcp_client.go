package main

import (
	"bufio"
	"fmt"
	"net"
)

func TcpConnect(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(conn)
	go func() {
		for {
			writer.Write([]byte("nihao"))
		}
	}()
	reader := bufio.NewReader(conn)
	for {
		// lenData, err := reader.Peek(2)
		// if err != nil {
		// 	fmt.Println("read from server failed,error:", err)
		// 	continue
		// }
		// fmt.Printf("收到server端发送的数据：%x\n", lenData)
		// data := make([]byte, lenData[1])
		data := make([]byte, 2)
		n, err := reader.Read(data)
		if err != nil {
			fmt.Println("read from server failed,error:", err)
			continue
		}
		fmt.Printf("收到server端发送的数据：%x\n", string(data[:n]))
	}
}

func main() {
	TcpConnect(":77")
}
