package webidl

type DOMException interface {
	isDOMException()
	Name() string
	Message() string
	Code() uint16
	error
}

const (
	DOMExceptionIndexSizeErr             uint16 = 1
	DOMExceptionDOMStringSizeErr         uint16 = 2
	DOMExceptionHierarchyRequestErr      uint16 = 3
	DOMExceptionWrongDocumentErr         uint16 = 4
	DOMExceptionInvalidCharacterErr      uint16 = 5
	DOMExceptionNoModificationAllowedErr uint16 = 6
	DOMExceptionNotFoundErr              uint16 = 7
	DOMExceptionNotSupportedErr          uint16 = 8
	DOMExceptionInUseAttributeErr        uint16 = 9
	DOMExceptionInvalidStateErr          uint16 = 10
	DOMExceptionSyntaxErr                uint16 = 11
	DOMExceptionInvalidModificationErr   uint16 = 12
	DOMExceptionNamespaceErr             uint16 = 13
	DOMExceptionInvalidAccessErr         uint16 = 14
	DOMExceptionValidationErr            uint16 = 15
	DOMExceptionTypeMismatchErr          uint16 = 16
	DOMExceptionSecurityErr              uint16 = 17
	DOMExceptionNetworkErr               uint16 = 18
	DOMExceptionAbortErr                 uint16 = 19
	DOMExceptionURLMismatchErr           uint16 = 20
	DOMExceptionQuotaExceededErr         uint16 = 21
	DOMExceptionTimeoutErr               uint16 = 22
	DOMExceptionInvalidNodeTypeErr       uint16 = 23
	DOMExceptionDataCloneErr             uint16 = 24
)

func NewDOMException(message *string, name *string) DOMException {
	return newDOMException(message, name)
}
