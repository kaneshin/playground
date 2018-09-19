package irrational

import (
	"math"
	"testing"
)

const (
	eps = 1e-2
)

func Test_Pi(t *testing.T) {
	tests := map[string]struct {
		result float
		expect float
	}{
		"Pi": {Pi, math.Pi},
	}

	for k, tt := range tests {
		tt := tt
		t.Run(k, func(t *testing.T) {
			if !(tt.expect-eps < tt.result && tt.result < tt.expect+eps) {
				t.Errorf("want %v (±%v) but got %v", tt.expect, eps, tt.result)
			}
			t.Logf("result %v (accuracy %v)", tt.result, math.Abs(float64(tt.result-tt.expect)))
		})
	}
}
