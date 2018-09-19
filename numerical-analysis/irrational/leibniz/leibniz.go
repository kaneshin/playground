package leibniz

import (
	"github.com/kaneshin/playground/numerical-analysis/internal"
	"github.com/kaneshin/playground/numerical-analysis/irrational"
)

type (
	float = internal.Float
)

func init() {
	irrational.Pi = leibnizFormula()
}

func leibnizFormula() float {
	const n = 1000
	p := float(0.)

	for i := 0; i < n; i += 2 {
		p += 1. / (2.*float(i) + 1.)
	}
	for i := 1; i < n; i += 2 {
		p -= 1. / (2.*float(i) + 1.)
	}
	return 4. * p
}
