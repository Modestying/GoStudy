package main

import (
	"fmt"
	"net"
)

func tcpServer(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		fmt.Println("get tcp")
		conn.Close()
	}
}

// ./main.exe 1 2 "third" --port=15
func main() {
	fmt.Println(getIntranetIP())
	// 监听TCP 服务端口
	listener, err := net.Listen("tcp", "0.0.0.0:8503")
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
