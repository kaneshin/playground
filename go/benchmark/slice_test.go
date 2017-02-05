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

func mergeAAndB(a, b []int) []int {
	ret := make([]int, 0, len(a)+len(b))
	const step = 2

	i := 0
	for _, b := range b {
		left := i * step
		right := (i + 1) * step
		ret = append(ret, a[left:right]...)
		ret = append(ret, b)
		_ = b
		i++
	}

	return append(ret, a[i*step:]...)
}

var c []int

func BenchmarkMergeAAndB(b *testing.B) {
	as := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	bs := []int{
		91, 92, 93,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = mergeAAndB(as, bs)
	}
}
