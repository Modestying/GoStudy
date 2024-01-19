package main

import (
	"fmt"
	"reflect"
)

type A int

type B = int

func main() {
	var a A
	var b B
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

}
