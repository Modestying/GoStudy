package main

import (
	"fmt"
)

func RangeArrary() {
	// 数组遍历
	// 1
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		if i == 0 {
			arr[1] = 100
			arr[2] = 200
			fmt.Println(arr)
		}
		arr[i] = v + 100
	}
	fmt.Println(arr)
}

func RangeSlice() {
	// 1,100,200   101,200,300
	data := []int{1, 2, 3}
	for i, v := range data {
		if i == 0 {
			data[1] = 100
			data[2] = 200
			fmt.Println(data)
		}

		data[i] = v + 100
	}
	fmt.Println(data)
}

func main() {
	//RangeArrary()
	//RangeSlice()
	ma := make(map[int]string)
	d := ma[1]
	if d == "" {
		fmt.Println("xx")
	}
	fmt.Println(d)
}
