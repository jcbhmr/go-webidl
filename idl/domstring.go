package idl

type DOMString struct {
	inner string
}

func (d *DOMString) Lift(s string) {
	d.inner = s
}

func (d DOMString) Lower() string {
	return d.inner
}
