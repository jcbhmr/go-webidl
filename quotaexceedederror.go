package webidl

type QuotaExceededError interface {
	isQuotaExceededError()
	DOMException
	Quota() *float64
	Requested() *float64
}

type QuotaExceededErrorOptions struct {
	Quota     *float64
	Requested *float64
}

func NewQuotaExceededError(message *string, options *QuotaExceededErrorOptions) (QuotaExceededError, error) {
	return newQuotaExceededError(message, options)
}
