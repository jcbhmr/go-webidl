package idl

type ArrayBuffer struct {
	inner []byte
}

func (a *ArrayBuffer) Lift(b []byte) {
	a.inner = b
}

func (a *ArrayBuffer) Lower() []byte {
	return a.inner
}

type SharedArrayBuffer struct {
	inner []byte
}

func (s *SharedArrayBuffer) Check() {
	if s.inner == nil {
		panic("SharedArrayBuffer: inner buffer is nil")
	}
}

func (s *SharedArrayBuffer) Lift(b []byte) {
	s.inner = b
}

func (s *SharedArrayBuffer) Lower() []byte {
	s.Check()
	return s.inner
}

type DataView struct {
	buffer BufferType
}

func (d DataView) Check() {
	CheckBufferType(d.buffer)
}

func (d *DataView) Lift(b BufferType) {
	CheckBufferType(b)
	d.buffer = b
}

func (d DataView) Lower() BufferType {
	d.Check()
	return d.buffer
}

// The buffer types are ArrayBuffer and SharedArrayBuffer.
//
// ArrayBuffer | SharedArrayBuffer
type BufferType = any

func CheckBufferType(v any) {
	switch v := v.(type) {
	case *ArrayBuffer:
		if v == nil {
			panic("TypeError: ArrayBuffer cannot be nil")
		}
	case *SharedArrayBuffer:
		if v == nil {
			panic("TypeError: SharedArrayBuffer cannot be nil")
		}
	default:
		panic("TypeError: expected ArrayBuffer or SharedArrayBuffer")
	}
}
