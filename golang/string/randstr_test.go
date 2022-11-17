package string

import "testing"

func BenchmarkRandStr(b *testing.B) {
	b.Run("RandStr1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RandStr(5)
		}
	})
	b.Run("RandStr2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RandStr2(5)
		}
	})
}
