package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Ptr *int
type Demo struct {
}

func (d *Demo) Func1(a int) {

}

func (d *Demo) Func2(a *int) int {
	return 0
}
func Test() {
	var wg Demo
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		for j := 1; j < method.Type.NumIn(); j++ {
			fmt.Println(method.Type.In(j).Name())
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(), method.Name, strings.Join(argv, ""), strings.Join(returns, ""))
	}
}

func TestReflect() {
	var wg []int
	//åvar wg map[string]int
	fmt.Println(reflect.TypeOf(wg))        //直接调用，返回的是wg的类型，是一个waitgroup 的指针
	fmt.Println(reflect.TypeOf(wg).Elem()) // 返回wg的类型指向的值
	fmt.Println(reflect.TypeOf(wg).Kind())
	fmt.Println(reflect.TypeOf(wg).Elem().Kind())
}

type Demo1 struct {
}

func TestDemo(d interface{}) {
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(d).Name())
	fmt.Println(reflect.Indirect(reflect.ValueOf(d)).String())

}
func main() {
	TestDemo(&Demo1{})
	TestDemo(Demo1{})
}
