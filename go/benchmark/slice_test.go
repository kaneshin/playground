package benchmark

import (
	"testing"
)

func BenchmarkSlice(b *testing.B) {
	const n = 100000

	b.Run("None", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			var dst []int
			b.StartTimer()
			for i := 0; i < n; i++ {
				dst = append(dst, i)
			}
		}
	})

	b.Run("Len", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			dst := make([]int, n)
			b.StartTimer()
			for i := 0; i < n; i++ {
				dst[i] = i
			}
		}
	})

	b.Run("Cap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			dst := make([]int, 0, n)
			b.StartTimer()
			for i := 0; i < n; i++ {
				dst = append(dst, i)
			}
		}
	})

	b.Run("Len-Cap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			dst := make([]int, n, n)
			b.StartTimer()
			for i := 0; i < n; i++ {
				dst[i] = i
			}
		}
	})
}
