package main

import (
	"flag"
	"fmt"
)

var port int

func init() {
	flag.IntVar(&port, "port", 1234, "Just for demo")
}

// ./main.exe 1 2 "third" --port=15
func main() {
	flag.Parse()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("port=", port)
}
