package webidl

import (
	"sync"
	"syscall/js"
)

var jsReflect = sync.OnceValue(func() js.Value {
	return js.Global().Get("Reflect")
})

func recoverJSError(errp *error) {
	r := recover()
	if r == nil {
		return
	}
	if jsErr, ok := r.(js.Error); ok {
		prototype := jsReflect().Call("getPrototypeOf", jsErr.Value)
		if prototype.Equal(jsEvalError()) {
			*errp = evalError{value: jsErr.Value}
		} else if prototype.Equal(jsRangeError()) {
			*errp = rangeError{value: jsErr.Value}
		} else if prototype.Equal(jsReferenceError()) {
			*errp = referenceError{value: jsErr.Value}
		} else if prototype.Equal(jsTypeError()) {
			*errp = typeError{value: jsErr.Value}
		} else if prototype.Equal(jsURIError()) {
			*errp = uriError{value: jsErr.Value}
		} else if prototype.Equal(jsDOMException()) {
			*errp = domException{value: jsErr.Value}
		} else if prototype.Equal(jsQuotaExceededError()) {
			*errp = quotaExceededError{domException: domException{value: jsErr.Value}}
		} else {
			*errp = jsErr
		}
	} else {
		panic(r)
	}
}
