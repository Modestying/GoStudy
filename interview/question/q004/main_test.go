package main

import (
	"reflect"
	"testing"
)

func TestIsSameStr(t *testing.T) {
	type TestStr struct {
		Str1  string
		Str2  string
		Equal bool
	}

	data := []TestStr{
		{
			Str1:  "sasd",
			Str2:  "sads",
			Equal: true,
		},
	}
	for _, val := range data {
		if !reflect.DeepEqual(val.Equal, IsSameStr(val.Str1, val.Str2)) {
			t.Error("wrong case:", val)
		} else {
			t.Log("pass case:", val)
		}
	}
}
