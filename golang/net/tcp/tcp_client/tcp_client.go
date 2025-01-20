package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":77")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	log.Println("close ok")

	conn.Write([]byte{0x01, 0x02})
	time.Sleep(time.Second)
	conn.Close()

}
