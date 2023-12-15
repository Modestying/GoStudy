package main

type Peole interface {
	SayHello()
}

type Student struct {
	Name string
}

func (s *Student) SayHello() {
	println("hello")
}

func main() {
	var stu Peole = &Student{}
	stu.SayHello()
}
