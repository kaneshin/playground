package shuffle

import (
	"math/rand"
	"time"
)

type Interface interface {
	Seed() int64
	Len() int
	Swap(i, j int)
}

func Shuffle(data Interface) {
	rand.Seed(data.Seed())
	n := data.Len()
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data.Swap(i, j)
	}
}

type IntSlice []int

func (p IntSlice) Seed() int64   { return time.Now().UnixNano() }
func (p IntSlice) Len() int      { return len(p) }
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p IntSlice) Shuffle()      { Shuffle(p) }

type Int64Slice []int64

func (p Int64Slice) Seed() int64   { return time.Now().UnixNano() }
func (p Int64Slice) Len() int      { return len(p) }
func (p Int64Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Int64Slice) Shuffle()      { Shuffle(p) }

type Float64Slice []float64

func (p Float64Slice) Seed() int64   { return time.Now().UnixNano() }
func (p Float64Slice) Len() int      { return len(p) }
func (p Float64Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Float64Slice) Shuffle()      { Shuffle(p) }
