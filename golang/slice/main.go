package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[1:3]
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(sli))
}
