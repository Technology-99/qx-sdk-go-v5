package qxErrs

import "errors"

var (
	ErrRemoteNotHealthy = errors.New("remote not healthy")
	ErrNotReady         = errors.New("qx sdk not ready")
	ErrNotLogined       = errors.New("qx sdk not logined")
	ErrMaxErrTimes      = errors.New("network error, max retry times")
)
