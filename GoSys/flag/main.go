package main

import (
	"flag"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("Read from tcp server failed,err:", err)
			break
		}
		data := string(buf[:n])
		fmt.Printf("Recived from client,data:%s\n", data)
	}
}

var port int
var testBool bool
func init() {
	flag.IntVar(&port, "port", 1234, "tcp port")
	flag.BoolVar(&testBool,"open",false,"Open tcp?")
}
//-flag //只支持bool类型
//-flag=x
//-flag x //只支持非bool类型
// ./main.exe 1 2 "third" --port=15

/*
$ ./main.exe -port 56 -open
args=[], num=0
port= 56
Open  true
*/

/*
$ ./main.exe -port 56
args=[], num=0
port= 56
Open  false
*/

func main() {
	flag.Parse()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("port=", port)
	fmt.Println("Open ",testBool)
}
