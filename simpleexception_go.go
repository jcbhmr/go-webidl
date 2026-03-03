//go:build !js && !nodeapi

package webidl

type evalError string

var _ EvalError = evalError("")

func (e evalError) isEvalError() {}

func (e evalError) Error() string {
	return string(e)
}

func newEvalError(message *string) EvalError {
	var goMessage string
	if message != nil {
		goMessage = *message
	}
	return evalError(goMessage)
}

type rangeError string

var _ RangeError = rangeError("")

func (r rangeError) isRangeError() {}

func (r rangeError) Error() string {
	return string(r)
}

func newRangeError(message *string) RangeError {
	var goMessage string
	if message != nil {
		goMessage = *message
	}
	return rangeError(goMessage)
}

type referenceError string

var _ ReferenceError = referenceError("")

func (r referenceError) isReferenceError() {}

func (r referenceError) Error() string {
	return string(r)
}

func newReferenceError(message *string) ReferenceError {
	var goMessage string
	if message != nil {
		goMessage = *message
	}
	return referenceError(goMessage)
}

type typeError string

var _ TypeError = typeError("")

func (t typeError) isTypeError() {}

func (t typeError) Error() string {
	return string(t)
}

func newTypeError(message *string) TypeError {
	var goMessage string
	if message != nil {
		goMessage = *message
	}
	return typeError(goMessage)
}

type uriError string

var _ URIError = uriError("")

func (u uriError) isURIError() {}

func (u uriError) Error() string {
	return string(u)
}

func newURIError(message *string) URIError {
	var goMessage string
	if message != nil {
		goMessage = *message
	}
	return uriError(goMessage)
}
