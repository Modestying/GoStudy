package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(RandStr2(5))
		fmt.Println("---------------------------------")
	}

}

//RandStr
func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Int31()%62])
	}
	return string(result)
}

func RandStr2(len int) string {
	result := make([]byte, len)
	time.Sleep(time.Millisecond * 1)
	rand.Seed(time.Now().UnixNano())
	rand.Read(result)
	return hex.EncodeToString(result)
}
