//go:build !js && !nodeapi

package webidl

type domException struct {
	name    string
	message string
}

var _ DOMException = (*domException)(nil)

func (d *domException) isDOMException() {}

func (d *domException) Name() string {
	return d.name
}

func (d *domException) Message() string {
	return d.message
}

var errorNamesToCodes = map[string]uint16{
	"IndexSizeError":             DOMExceptionIndexSizeErr,
	"HierarchyRequestError":      DOMExceptionHierarchyRequestErr,
	"WrongDocumentError":         DOMExceptionWrongDocumentErr,
	"InvalidCharacterError":      DOMExceptionInvalidCharacterErr,
	"NoModificationAllowedError": DOMExceptionNoModificationAllowedErr,
	"NotFoundError":              DOMExceptionNotFoundErr,
	"NotSupportedError":          DOMExceptionNotSupportedErr,
	"InUseAttributeError":        DOMExceptionInUseAttributeErr,
	"InvalidStateError":          DOMExceptionInvalidStateErr,
	"SyntaxError":                DOMExceptionSyntaxErr,
	"InvalidModificationError":   DOMExceptionInvalidModificationErr,
	"NamespaceError":             DOMExceptionNamespaceErr,
	"InvalidAccessError":         DOMExceptionInvalidAccessErr,
	"TypeMismatchError":          DOMExceptionTypeMismatchErr,
	"SecurityError":              DOMExceptionSecurityErr,
	"NetworkError":               DOMExceptionNetworkErr,
	"AbortError":                 DOMExceptionAbortErr,
	"URLMismatchError":           DOMExceptionURLMismatchErr,
	"QuotaExceededError":         DOMExceptionQuotaExceededErr,
	"TimeoutError":               DOMExceptionTimeoutErr,
	"InvalidNodeTypeError":       DOMExceptionInvalidNodeTypeErr,
	"DataCloneError":             DOMExceptionDataCloneErr,
}

func (d *domException) Code() uint16 {
	code, ok := errorNamesToCodes[d.name]
	if !ok {
		return 0
	}
	return code
}

func (d *domException) Error() string {
	if d.message != "" {
		return d.name + ": " + d.message
	} else {
		return d.name
	}
}

func newDOMException(message *string, name *string) DOMException {
	goMessage := ""
	if message != nil {
		goMessage = *message
	}
	goName := "Error"
	if name != nil {
		goName = *name
	}
	return &domException{name: goName, message: goMessage}
}
