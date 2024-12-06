package main

type People interface {
	SayHello()
}

type Student struct {
	Name string
}

func (s *Student) SayHello() {
	println("hello")
}

func main() {
	var stu People = &Student{}
	stu.SayHello()
}
