package webidl

import (
	"sync"
	"syscall/js"
)

var jsEvalError = sync.OnceValue(func() js.Value {
	return js.Global().Get("EvalError")
})
var jsRangeError = sync.OnceValue(func() js.Value {
	return js.Global().Get("RangeError")
})
var jsReferenceError = sync.OnceValue(func() js.Value {
	return js.Global().Get("ReferenceError")
})
var jsTypeError = sync.OnceValue(func() js.Value {
	return js.Global().Get("TypeError")
})
var jsURIError = sync.OnceValue(func() js.Value {
	return js.Global().Get("URIError")
})

type evalError struct {
	value js.Value
}

var _ EvalError = evalError{}

func (e evalError) isEvalError() {}

func (e evalError) Error() string {
	str := e.value.Call("toString")
	if str.Type() != js.TypeString {
		panic(&js.ValueError{Method: "Value.String", Type: e.value.Type()})
	}
	return str.String()
}

func newEvalError(message *string) EvalError {
	var jsMessage js.Value
	if message != nil {
		jsMessage = js.ValueOf(*message)
	}
	return evalError{value: jsEvalError().New(jsMessage)}
}

type rangeError struct {
	value js.Value
}

var _ RangeError = rangeError{}

func (r rangeError) isRangeError() {}

func (r rangeError) Error() string {
	str := r.value.Call("toString")
	if str.Type() != js.TypeString {
		panic(&js.ValueError{Method: "Value.String", Type: r.value.Type()})
	}
	return str.String()
}

func newRangeError(message *string) RangeError {
	var jsMessage js.Value
	if message != nil {
		jsMessage = js.ValueOf(*message)
	}
	return rangeError{value: jsRangeError().New(jsMessage)}
}

type referenceError struct {
	value js.Value
}

var _ ReferenceError = referenceError{}

func (r referenceError) isReferenceError() {}

func (r referenceError) Error() string {
	str := r.value.Call("toString")
	if str.Type() != js.TypeString {
		panic(&js.ValueError{Method: "Value.String", Type: r.value.Type()})
	}
	return str.String()
}

func newReferenceError(message *string) ReferenceError {
	var jsMessage js.Value
	if message != nil {
		jsMessage = js.ValueOf(*message)
	}
	return referenceError{value: jsReferenceError().New(jsMessage)}
}

type typeError struct {
	value js.Value
}

var _ TypeError = typeError{}

func (t typeError) isTypeError() {}

func (t typeError) Error() string {
	str := t.value.Call("toString")
	if str.Type() != js.TypeString {
		panic(&js.ValueError{Method: "Value.String", Type: t.value.Type()})
	}
	return str.String()
}

func newTypeError(message *string) TypeError {
	var jsMessage js.Value
	if message != nil {
		jsMessage = js.ValueOf(*message)
	}
	return typeError{value: jsTypeError().New(jsMessage)}
}

type uriError struct {
	value js.Value
}

var _ URIError = uriError{}

func (u uriError) isURIError() {}

func (u uriError) Error() string {
	str := u.value.Call("toString")
	if str.Type() != js.TypeString {
		panic(&js.ValueError{Method: "Value.String", Type: u.value.Type()})
	}
	return str.String()
}

func newURIError(message *string) URIError {
	var jsMessage js.Value
	if message != nil {
		jsMessage = js.ValueOf(*message)
	}
	return uriError{value: jsURIError().New(jsMessage)}
}
