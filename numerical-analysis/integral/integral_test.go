package integral

import (
	"math"
	"testing"

	"github.com/kaneshin/playground/numerical-analysis/internal"
)

const (
	eps = 1e-3
)

type impl struct {
	f    function
	a, b float
}

func (i impl) Func() function {
	return i.f
}

func (i impl) Interval() (float, float) {
	return i.a, i.b
}

func (i impl) Step() int {
	return 1e+3
}

func TestIntegral(t *testing.T) {
	tests := map[string]struct {
		data   Interface
		expect float
	}{
		// http://m.wolframalpha.com/input/?i=integral_0%5E1+%28x%5E2+-+x+-1+%29
		"∫[0,1] F1 dx": {impl{internal.F1, 0, 1}, -7. / 6.},
		// http://m.wolframalpha.com/input/?i=integral_-2%5E4+%28x%5E2+-+x+-1+%29
		"∫[-2,4] F1 dx": {impl{internal.F1, -2, 4}, 12.},
	}

	for k, tt := range tests {
		tt := tt
		t.Run(k, func(t *testing.T) {
			result := Integral(tt.data)
			if !(tt.expect-eps < result && result < tt.expect+eps) {
				t.Errorf("want %v (±%v) but got %v", tt.expect, eps, result)
			}
			t.Logf("result %v (accuracy %v)", result, math.Abs(float64(result-tt.expect)))
		})
	}
}
