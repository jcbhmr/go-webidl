package webidl_test

import (
	"testing"

	"go.jcbhmr.com/webidl"
)

func TestNewDOMException(t *testing.T) {
	d := webidl.NewDOMException(nil, nil)
	if d == nil {
		t.Fatal("expected non-nil DOMException")
	}
}

func TestDOMException_Error(t *testing.T) {
	d := webidl.NewDOMException(nil, nil)
	e := d.Error()
	if e == "" {
		t.Fatal("expected non-empty error string")
	}
}
