package main

import "fmt"


/*
负载因子：kv数量/桶数量 6.5

*/
type Student struct {
	Age int
}

func main() {
	kv := map[string]string{"menglu": "xx"}
	x := kv["s"]
	if x == "" {
		fmt.Println("ss")
	}
}
