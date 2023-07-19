package main

import (
	"net/http"
	"sync"
)

// sync.pool

type MyStruct struct {
	http.Request
	http.Response
	a http.Request
	b http.Request
	c http.Request
	d http.Request
}

func NewDemo() *MyStruct {
	return &MyStruct{}
}

var (
	demoPool = sync.Pool{
		New: func() interface{} {
			return &MyStruct{}
		},
	}
)

func NewDemoWithPool() *MyStruct {
	return demoPool.Get().(*MyStruct)
}
