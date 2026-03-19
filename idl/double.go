package idl

import "math"

type Double struct {
	inner float64
}

func (d *Double) Lift(f float64) {
	if math.IsInf(f, 0) || math.IsNaN(f) {
		panic("TypeError: double cannot be NaN or Inf")
	}
	d.inner = f
}

func (d Double) Lower() float64 {
	return d.inner
}
