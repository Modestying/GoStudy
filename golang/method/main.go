package main

import "fmt"

type A struct {
}

func (a A) Name() {
	fmt.Println("xx")
}

func main() {
	a := A{}
	a.Name()
	A.Name(a)
}
