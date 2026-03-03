package webidl

import (
	"sync"
	"syscall/js"
)

type domException struct {
	value js.Value
}

var jsDOMException = sync.OnceValue(func() js.Value {
	return js.Global().Get("DOMException")
})

var _ DOMException = domException{}

func (d domException) isDOMException() {}

func (d domException) Name() string {
	name := d.value.Get("name")
	if name.Type() != js.TypeString {
		panic(&js.ValueError{Method: "Value.String", Type: d.value.Type()})
	}
	return name.String()
}

func (d domException) Message() string {
	message := d.value.Get("message")
	if message.Type() != js.TypeString {
		panic(&js.ValueError{Method: "Value.String", Type: d.value.Type()})
	}
	return message.String()
}

func (d domException) Code() uint16 {
	return uint16(d.value.Get("code").Int())
}

func (d domException) Error() string {
	str := d.value.Call("toString")
	if str.Type() != js.TypeString {
		panic(&js.ValueError{Method: "Value.String", Type: d.value.Type()})
	}
	return str.String()
}

func newDOMException(message *string, name *string) DOMException {
	var jsMessage js.Value
	if message != nil {
		jsMessage = js.ValueOf(*message)
	}
	var jsName js.Value
	if name != nil {
		jsName = js.ValueOf(*name)
	}
	return domException{value: jsDOMException().New(jsMessage, jsName)}
}
