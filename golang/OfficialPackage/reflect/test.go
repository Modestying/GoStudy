package main

import (
	"fmt"
	"reflect"
)

func dd() {
	servo := &Servo{data: 1}
	value := reflect.ValueOf(servo)
	fun := value.MethodByName("PanLeft")
	if !fun.IsValid() {
		fmt.Println("invalid")
	} else {
		fun.Call([]reflect.Value{})
	}

	//arg := make([]reflect.Value, 2)
	//arg[0] = reflect.ValueOf(1)
	//arg[1] = reflect.ValueOf("dd")
	//reflect.ValueOf(servo).MethodByName("WithVal").Call(arg)

	reflect.ValueOf(servo).MethodByName("WithVal2").Call([]reflect.Value{reflect.ValueOf(D{
		name: "11",
		num:  12,
	})})

	f:=reflect.ValueOf(servo).MethodByName("WithVal")
	fmt.Println(f.Type().NumIn()) //获取函数参数个数
	fmt.Println(reflect.ValueOf(f))
}

