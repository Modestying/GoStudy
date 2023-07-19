package main

import "fmt"

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
