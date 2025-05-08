package qxErrs

import "errors"

var (
	ErrRemoteNotHealthy        = errors.New("remote not healthy")
	ErrCannotCreateConfigDir   = errors.New("qx sdk can't create qx config directory")
	ErrConfigDirIsNotDirectory = errors.New("qx sdk qx config directory is not a directory")
	ErrAuthFileIsNotFile       = errors.New("qx sdk qx auth file is not file")
	ErrNotReady                = errors.New("qx sdk not ready")
	ErrNotLogined              = errors.New("qx sdk not logined")
	ErrMaxErrTimes             = errors.New("network error, max retry times")
)
