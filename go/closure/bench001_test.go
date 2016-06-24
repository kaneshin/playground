package bench

import (
	"testing"
)

func Normal(n int) string {
	orz := map[int]string{
		1: "Golang",
		2: "Clojure",
		3: "Swift",
	}
	return orz[n]
}

var Closure = func() func(n int) string {
	orz := map[int]string{
		1: "Golang",
		2: "Clojure",
		3: "Swift",
	}
	return func(n int) string {
		return orz[n]
	}
}()

func BenchmarkNormal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Normal(n%3 + 1)
	}
}

func BenchmarkClosure(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Closure(n%3 + 1)
	}
}
