package maths

import (
	"math"
	nr "maths/internal/maths"
	"testing"
)

func TestNewtonRaphsonWithAbsoluteErrorStoppingCriteria(t *testing.T) {
	epsilon := 0.000001
	newtonRaphson := nr.NewNewtonRaphson(
		func(x float64) float64 {
			return math.Pow(x, 5) - 2*math.Pow(x, 3) - math.Log(x)
		},
		func(x float64) float64 {
			return 5*math.Pow(x, 4) - 6*math.Pow(x, 2) - 1/x
		},
		1,
		epsilon,
		nr.AbsoluteError,
	)
	got := newtonRaphson.Solve()
	approxValue := 0.649233541
	want := approxValue
	if math.Abs(want-got) >= epsilon {
		t.Errorf("got %.20f, want %.20f", got, want)
	}
}
