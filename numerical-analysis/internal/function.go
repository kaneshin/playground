package internal

type Float float64
type Function func(...Float) Float

func F1(x ...Float) Float {
	if len(x) < 1 {
		return 0.
	}
	x0 := x[0]
	return x0*x0 - x0 - 1
}

func F2(x ...Float) Float {
	if len(x) < 2 {
		return 0.
	}
	x0 := x[0]
	x1 := x[1]
	return x0*x1 - x0 - x1 - 1
}
