package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func tcpServer(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		reader := bufio.NewReader(conn)
		buf := make([]byte, 2048)
		if _, err := io.ReadAtLeast(reader, buf, 4); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", buf)
		conn.Close()
	}
}

// ./main.exe 1 2 "third" --port=15
func main() {
	fmt.Println(getIntranetIP())
	// 监听TCP 服务端口
	listener, err := net.Listen("tcp", "127.0.0.1:8503")
	if err != nil {
		fmt.Println("Listen tcp server failed,err:", err)
		return
	}
	defer listener.Close()
	tcpServer(listener)
}

func getIntranetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
