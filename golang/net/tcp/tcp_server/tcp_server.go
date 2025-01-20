package main

import (
	"fmt"
	"net"
	"time"
)

var data = []byte{0x01, 0x02}

func process(conn net.Conn) {
	defer func() {
		fmt.Printf("Conn close:%s", conn.RemoteAddr())
	}()
	for {
		time.Sleep(time.Second * 2)
		b := make([]byte, 2000)
		_, err := conn.Read(b)
		if err != nil {
			fmt.Println("read from client failed,error:", err)
			conn.Close()
			return
		}
	}
}

/*
615455 before send
615628 send data

615632 after send
615682 send rst
*/

// StartTcpServer 启动httpServer
func StartTcpServer(address string) {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Listen failed,error:", err)
		return
	} else {
		fmt.Println("success listen!")
	}
	fmt.Println("waiting for client")
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed,error:", err)
			continue
		} else {
			fmt.Println("success accept!")
		}
		go process(conn) //启动一个goroutine处理连接
	}
}

func main() {
	StartTcpServer(":77")
}
