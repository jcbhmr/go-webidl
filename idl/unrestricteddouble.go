package idl

type UnrestrictedDouble struct {
	inner float64
}

func (d *UnrestrictedDouble) Lift(f float64) {
	d.inner = f
}

func (d UnrestrictedDouble) Lower() float64 {
	return d.inner
}
