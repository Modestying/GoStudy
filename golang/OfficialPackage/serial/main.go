package main

import (
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"io"
)

var port string
var baudRate int

func init() {
	flag.StringVar(&port, "port", "ttyPS1", "串口")
	flag.IntVar(&baudRate, "rate", 115200, "串口波特率")
}
func main() {
	flag.Parse()
	fmt.Println("port=", port)
	fmt.Println("rate=", baudRate)
	serialPort, err := serial.OpenPort(&serial.Config{Name: port, Baud: baudRate, ReadTimeout: 500})
	if err != nil {
		fmt.Println("Open serial port failed ", err)
	}
	defer func() {
		serialPort.Close()
		fmt.Println("close port")
	}()

	Read(serialPort)
}

func Read(c io.ReadWriteCloser) {
	for {
		data := make([]byte, 1024)
		if n, err := c.Read(data); err == nil && n != 0 {
			fmt.Printf("Read len:%d  %X\n", n, data[:n])
			c.Write(data[:n])
		}
	}
}
