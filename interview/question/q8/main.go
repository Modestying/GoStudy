package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for ket, stu := range stus {
		m[stu.Name] = &stus[ket]
	}
	for _, val := range m {
		fmt.Println(val)
	}
}

func main() {
	pase_student()

}
