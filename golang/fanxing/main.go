package main

import "fmt"

func sum[V ~int | ~float64 | ~string](vs ...V) V {
	var r V
	for _, v := range vs {
		r += v
	}
	return r
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1.1, 2.2, 3.3))
}
