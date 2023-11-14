package main

import "testing"

func BenchmarkSumWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumWithDefer()
	}
}

func BenchmarkSumWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumWithoutDefer()
	}
}
