package string_test

import (
	"testing"

	"github.com/Modestying/GoStudy/golang/string"
)

func BenchmarkRandStr(b *testing.B) {
	b.Run("RandStr1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			string.RandStr(5)
		}
	})
	b.Run("RandStr2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			string.RandStr2(5)
		}
	})
}
