package main

import (
	"fmt"
	"testing"
)

func BenchmarkNextVal(b *testing.B) {
	inst, _ := NewIdWoker(0, 0)
	var i int64
	for i = 0; i < sequenceMax; i++ {
		fmt.Println(inst.NextId())
	}
}
