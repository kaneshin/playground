package benchmark

import (
	"fmt"
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	b.Run("Bar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprint("%d", i)
		}
	})
}
