package webidl

type EvalError interface {
	isEvalError()
	error
}

func NewEvalError(message *string) EvalError {
	return newEvalError(message)
}

type RangeError interface {
	isRangeError()
	error
}

func NewRangeError(message *string) RangeError {
	return newRangeError(message)
}

type ReferenceError interface {
	isReferenceError()
	error
}

func NewReferenceError(message *string) ReferenceError {
	return newReferenceError(message)
}

type TypeError interface {
	isTypeError()
	error
}

func NewTypeError(message *string) TypeError {
	return newTypeError(message)
}

type URIError interface {
	isURIError()
	error
}

func NewURIError(message *string) URIError {
	return newURIError(message)
}
