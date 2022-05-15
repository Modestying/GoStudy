package main

import "fmt"

type A struct {
	Name string
}

func (a *A) Do() {
	fmt.Println("A* do")
	a.Name = "a+"
}

type B struct {
	Name string
}

func (b B) Do() {
	fmt.Println("B do")
	b.Name = "b+"
}

func main() {
	a := A{Name: "a"}
	a.Do()
	fmt.Println(a.Name)
	a2 := &A{}
	a2.Do()
	fmt.Println(a2.Name)

	b := B{}
	b.Do()
	fmt.Println(b.Name)
	b2 := &B{}
	b2.Do()
	fmt.Println(b2.Name)
}
