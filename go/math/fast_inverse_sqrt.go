package main

import (
	"fmt"
	"unsafe"
)

func FastInverseSqrt(x float32) float32 {
	const threehalfs = float32(1.5)

	x2 := x * float32(0.5)
	y := x
	i := *(*int32)(unsafe.Pointer(&y))
	i = 0x5f3759df - i>>1
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (threehalfs - (x2 * y * y))
	return y
}

func main() {
	fmt.Printf("Fast inverse sqrt: %f\n", 1.0/FastInverseSqrt(2))
}
