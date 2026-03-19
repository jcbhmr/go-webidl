package idl

type Undefined struct{}

func (u *Undefined) Lift(v any) {}

func (u Undefined) Lower() struct{} {
	return struct{}{}
}
