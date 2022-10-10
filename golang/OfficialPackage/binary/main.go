package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{0x4e, 0x05, 0xfd, 0xff}
	binary.BigEndian.PutUint16(data[:2], uint16(1))
	fmt.Printf("%x\n", data)
}
