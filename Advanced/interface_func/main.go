package main

import "fmt"

type Foo struct {
	v int
}

func MakeFoo(n *int) Foo {
	print(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	print(*n)
}

func main() {
	var x = 1
	var p = &x
	c := 2

	defer func(p *int) {
		MakeFoo(p).Bar(p) // line 19
		fmt.Println(c)
	}(p)
	x = 2
	p = new(int) // line 21
	MakeFoo(p)
}
