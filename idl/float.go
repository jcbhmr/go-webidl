package idl

import "math"

type Float struct {
	inner float32
}

func (f *Float) Lift(v float32) {
	if math.IsInf(float64(v), 0) || math.IsNaN(float64(v)) {
		panic("TypeError: float cannot be NaN or Inf")
	}
	f.inner = v
}

func (f Float) Lower() float32 {
	return f.inner
}
