//go:build !js && !nodeapi

package webidl

type quotaExceededError struct {
	domException
	quota     *float64
	requested *float64
}

var _ QuotaExceededError = (*quotaExceededError)(nil)

func (q *quotaExceededError) isQuotaExceededError() {}

func (q *quotaExceededError) Quota() *float64 {
	return q.quota
}

func (q *quotaExceededError) Requested() *float64 {
	return q.requested
}

func newQuotaExceededError(message *string, options *QuotaExceededErrorOptions) (QuotaExceededError, error) {
	var goMessage string
	if message != nil {
		goMessage = *message
	}
	var goQuota *float64
	var goRequested *float64
	if options != nil {
		if options.Quota != nil {
			if *options.Quota < 0 {
				return nil, NewRangeError(new2("Quota cannot be negative"))
			}
			goQuota = options.Quota
		}
		if options.Requested != nil {
			goRequested = options.Requested
		}
	}
	return &quotaExceededError{
		domException: domException{name: "QuotaExceededError", message: goMessage},
		quota:        goQuota,
		requested:    goRequested,
	}, nil
}
