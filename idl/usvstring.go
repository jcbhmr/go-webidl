package idl

import (
	"iter"
	"strings"
	"unicode"
)

type USVString struct {
	utf8 string
}

func (u *USVString) Lift(s string) {
	u.utf8 = strings.ToValidUTF8(s, string(unicode.ReplacementChar))
}

func (u USVString) Lower() string {
	return u.utf8
}

func (u USVString) ScalarValues() iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range u.utf8 {
			if !yield(i, r) {
				return
			}
		}
	}
}
