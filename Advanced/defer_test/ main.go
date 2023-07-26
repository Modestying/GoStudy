package defer_test

import (
	"fmt"
	"runtime"
)

type Foo struct {
	v int
}

func MakeFoo(n *int, tag string) Foo {
	fmt.Println("tag: ", tag)
	fmt.Println(*n)
	return Foo{}
}

func (Foo) Bar(n *int) Zoo {
	fmt.Println("bar ", *n)
	return Zoo{}
}

func (Foo) Bar2(n *int) Zoo {
	fmt.Println("bar2 ", *n)
	return Zoo{}
}

type Zoo struct {
}

func (Zoo) ZooBar(n *int) {
	fmt.Println("zoo bar ", *n)
}

func (Zoo) ZooBar2(n *int) {
	fmt.Println("zoo bar2 ", *n)
}

func main() {
	var x = 1
	var p = &x
	defer MakeFoo(p, "defer tag ").Bar(p).ZooBar(p)
	defer MakeFoo(p, "defer tag ").Bar2(p).ZooBar2(p)

	bs := make([]byte, 1024)
	runtime.Stack(bs, false)
	fmt.Printf("current stack info: %s", string(bs))

	//x = 2
	//p = new(int) // line 21 0
	//
	//MakeFoo(p, "main p ")
}
