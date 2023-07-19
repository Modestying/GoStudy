package main

import "testing"

func BenchmarkNewDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo := NewDemo()
		_ = demo
	}
}

func BenchmarkNewDemoWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo := NewDemoWithPool()
		demoPool.Put(demo)
	}
}
