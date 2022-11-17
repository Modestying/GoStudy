package string

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// RandStr
func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano())
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
