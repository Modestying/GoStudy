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

func SliceDemoA() {
	str1 := []string{"1", "2", "3"}
	str2 := make([]string, 3)
	copy(str2, str1)
	fmt.Println(copy(str2, str1))
	str2[1] = "update"
	fmt.Println("str1 ", str1)
	fmt.Println("str2", str2)

}

func GetLen(s []int) {
	fmt.Printf("%d %d %v\n", len(s), cap(s), s)
}

func UpdateSlice(demo []int) {
	demo[0] = 11
}

func ApppendSlice(demo []int) {
	demo = append(demo, 11)
}
func main() {
	demo := make([]int, 5, 10)
	GetLen(demo)
	demo = append(demo, 1)
	GetLen(demo)
	UpdateSlice(demo)
	GetLen(demo)
	UpdateSlice(demo)
	GetLen(demo)

	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := a[1:5]
	s2 := a[7:]
	s1[0] = 11
	s2[0] = 12
	GetLen(s1)
	GetLen(s2)
	fmt.Println(a)
	s1 = append(s1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	s1[0] = 22
	GetLen(s1)
	fmt.Println(a)
}
