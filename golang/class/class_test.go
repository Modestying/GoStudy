package main

import "testing"

func TestDemo_GetName(t *testing.T) {
	demo := Demo{Name: "1"}

	demo.GetName()
	t.Log(demo.Name)

	// 对name的修改 未生效
	demo.GetNameOrigin()
	t.Log(demo.Name)
}
