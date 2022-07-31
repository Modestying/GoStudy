package dd

import (
	"testing"
)

func BelongArray(cmd []byte) bool {
	if cmd[0] == 0xa1 {
		return true
	}
	return false
}

func BelongHead(cmd byte) bool {
	if cmd == 0xa1 {
		return true
	}
	return false
}

func BenchmarkBelongArray(b *testing.B) {
	cmd := []byte{0xa1, 0xa2}
	b.Run("Head bench", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BelongHead(cmd[0])
		}
	})

	b.Run("Array bench", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BelongArray(cmd)
		}
	})
}
