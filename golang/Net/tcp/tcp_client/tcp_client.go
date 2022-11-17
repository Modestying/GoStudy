package tcp_client

import (
	"fmt"
	"net"
)

func TcpConnect(address string) {
	conn, err := net.Dial("tcp", address)
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
