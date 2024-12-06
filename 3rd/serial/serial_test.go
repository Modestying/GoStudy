package serial

import (
	"fmt"
	"testing"

	"github.com/tarm/serial"
)

func TestRead(t *testing.T) {
	//flag.Parse()
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
