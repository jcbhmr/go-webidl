package webidl

import (
	"sync"
	"syscall/js"
)

var jsQuotaExceededError = sync.OnceValue(func() js.Value {
	return js.Global().Get("QuotaExceededError")
})

type quotaExceededError struct {
	domException
}

var _ QuotaExceededError = quotaExceededError{}

func (q quotaExceededError) isQuotaExceededError() {}

func (q quotaExceededError) Quota() *float64 {
	quota := q.value.Get("quota")
	if quota.IsNull() || quota.IsUndefined() {
		return nil
	}
	goQuota := quota.Float()
	return &goQuota
}

func (q quotaExceededError) Requested() *float64 {
	requested := q.value.Get("requested")
	if requested.IsNull() || requested.IsUndefined() {
		return nil
	}
	goRequested := requested.Float()
	return &goRequested
}

func newQuotaExceededError(message *string, options *QuotaExceededErrorOptions) (_ QuotaExceededError, err error) {
	defer recoverJSError(&err)
	var jsMessage js.Value
	if message != nil {
		jsMessage = js.ValueOf(*message)
	}
	var jsOptions js.Value
	if options != nil {
		jsOptions = js.ValueOf(map[string]any{})
		if options.Quota != nil {
			jsOptions.Set("quota", *options.Quota)
		}
		if options.Requested != nil {
			jsOptions.Set("requested", *options.Requested)
		}
	}
	return quotaExceededError{domException: domException{value: jsQuotaExceededError().New(jsMessage, jsOptions)}}, nil
}
