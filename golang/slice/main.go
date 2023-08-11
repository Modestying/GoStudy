package main

import (
	"fmt"
	"reflect"
)

func SliceDemo() {
	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[1:3]
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(sli))
}

func main() {
	str1 := []string{"1", "2", "3"}
	str2 := make([]string, 3)
	copy(str2, str1)
	fmt.Println(copy(str2, str1))
	str2[1] = "update"
	fmt.Println("str1 ", str1)
	fmt.Println("str2", str2)

}
