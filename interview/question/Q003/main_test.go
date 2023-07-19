package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type TestStr struct {
		Src string
		Dst string
	}

	data := []TestStr{
		{
			Src: "123",
			Dst: "321",
		},
	}
	for _, val := range data {
		if res, ok := ReverseString(val.Src); ok {
			if !reflect.DeepEqual(val.Dst, res) {
				t.Error("wrong case:", val)
			} else {
				t.Log("pass case:", val)
			}
		}
	}
}
