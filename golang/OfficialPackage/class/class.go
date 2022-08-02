package main

type Demo struct {
	Name string
}

// struct 方法带*与不带*
//1. 带*,func中对内部成员变量的修改会保存
//2. instance和*instance都可以访问到GetNameOrigin和GetName方法

func (d *Demo) GetName() string {
	d.Name = "GetName"
	return d.Name
}

func (d Demo) GetNameOrigin() string {
	d.Name = "GetNameOrigin"
	return "Demo"
}
