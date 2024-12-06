package serial

import (
	"flag"
	"fmt"
	"io"
)

// TODO 更新串口使用
var port string
var baudRate int

func init() {
	flag.StringVar(&port, "port", "ttyPS1", "串口")
	flag.IntVar(&baudRate, "rate", 115200, "串口波特率")
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
