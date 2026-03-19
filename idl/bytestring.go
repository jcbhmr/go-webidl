package idl

import (
	"iter"
	"slices"
	"unsafe"
)

type ByteString struct {
	inner []byte
}

func (b *ByteString) Lift(s string) {
	b.inner = unsafe.Slice(unsafe.StringData(s), len(s))
}

func (b ByteString) Lower() string {
	return unsafe.String(unsafe.SliceData(b.inner), len(b.inner))
}

func (b ByteString) All() iter.Seq2[int, byte] {
	return slices.All(b.inner)
}
