package newton_raphson

import "math"

type NewtonRaphson struct {
    f            func(float64) float64
    fDx          func(float64) float64
    xn           float64
    xnPlus1      float64
    epsilon      float64
    stopCriteria uint8
}

func NewNewtonRaphson(f func(float64) float64, fDx func(float64) float64, xn, epsilon float64, stopCriteria uint8) NewtonRaphson {
    return NewtonRaphson{
        f:            f,
        fDx:          fDx,
        xn:           xn,
        xnPlus1:      0,
        epsilon:      epsilon,
        stopCriteria: stopCriteria,
    }
}

const (
    AbsoluteError uint8 = iota
    FuncAbsoluteValue
)

func (nr *NewtonRaphson) shouldStop() bool {
    value := func() float64 {
        if nr.stopCriteria == AbsoluteError {
            return nr.xnPlus1 - nr.xn
        } else if nr.stopCriteria == FuncAbsoluteValue {
            return nr.f(nr.xnPlus1)
        }
        return -1
    }()
    if value == -1 {
        panic("Invalid stop criteria")
    }
    return math.Abs(value) < nr.epsilon
}

func (nr *NewtonRaphson) iterate() float64 {
    return nr.xn - nr.f(nr.xn)/nr.fDx(nr.xn)
}

func (nr *NewtonRaphson) Solve() float64 {
    nr.xnPlus1 = nr.iterate()
    for !nr.shouldStop() {
        nr.xn = nr.xnPlus1
        nr.xnPlus1 = nr.iterate()
    }
    return nr.xnPlus1
}
