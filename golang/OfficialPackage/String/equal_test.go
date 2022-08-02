package main

import "testing"

func BenchmarkEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal("a", "a")
		Equal("a", "b")
	}
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compare("a", "a")
		Compare("a", "b")
	}
}

func BenchmarkEqualFold(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualFold("a", "b")
	}
}
