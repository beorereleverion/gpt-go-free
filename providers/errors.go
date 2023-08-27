package providers

import "errors"

var (
	errNoStreamSupport = errors.New("sorry, this provider has no stream support yet, please try another, or use NewCompletion func")
)
