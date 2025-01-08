package types

import "errors"

var (
	ErrNotReady    = errors.New("sdk not ready")
	ErrMaxErrTimes = errors.New("network error, max retry times")
)
