package benchmark

import (
	"reflect"
	"testing"
	"time"
)

func Test_Divide(t *testing.T) {
	const N = 1000000
	tmp := float64(1) / float64(time.Millisecond)
	for j := 0; j < N; j++ {
		f1 := int64(j) / int64(time.Millisecond)
		f2 := int64(float64(j) * tmp)
		if !reflect.DeepEqual(f1, f2) {
			t.Fatal("Should be the same values")
		}
	}
}

var f int64

func Benchmark_Divide(b *testing.B) {
	const N = 1000000
	b.Run("Divide", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < N; j++ {
				f = int64(j) / int64(time.Millisecond)
			}
		}
	})

	b.Run("Multiple", func(b *testing.B) {
		t := float64(1) / float64(time.Millisecond)
		for i := 0; i < b.N; i++ {
			for j := 0; j < N; j++ {
				f = int64(float64(j) * t)
			}
		}
	})
}
