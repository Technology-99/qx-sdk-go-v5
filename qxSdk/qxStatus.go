package qxSdk

const (
	SdkStatusNotReady = iota + 1
	SdkStatusRemoteUnHealthy
	SdkStatusRemoteHealthy
	SdkStatusSignMaxTimes
	SdkStatusLoginFailed
	SdkStatusLogined
	SdkStatusStartKeyExChange
	SdkStatusKeyExChanging
	SdkStatusReady
)
