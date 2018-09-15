package integral

import (
	"github.com/kaneshin/playground/numerical-analysis/internal"
)

type (
	function = internal.Function
	float    = internal.Float
)

type Interface interface {
	Func() function
	Interval() (float, float)
	Step() int
}

func Integral(data Interface) float {
	n := data.Step()
	if n <= 0 {
		return 0.
	}

	f := data.Func()
	a, b := data.Interval()
	if a > b {
		return -rectangularRule(f, b, a, n)
	}
	return rectangularRule(f, a, b, n)
}

func rectangularRule(f function, a, b float, n int) float {
	h := (b - a) / float(n)
	r := float(0.)

	// Evaluate by using midpoint formula
	x := a + h/2.
	for i := 0; i < n; i++ {
		r += f(x)
		x += h
	}
	return r * h
}

func trapezoidalRule(f function, a, b float, n int) float {
	panic("TODO: not implemented")
	return 0.
}

func simpsonRule(f function, a, b float, n int) float {
	panic("TODO: not implemented")
	return 0.
}
