package resp

type DefaultResponseBase struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestId string `json:"requestId"`
}

type QxV5ApisApiRefreshResp struct {
	QxV5ApisApiSignResp
}

type QxV5ApisApiSignResp struct {
	DefaultResponseBase
	Data QxV5ApisApiSignRespData `json:"data"`
}

type QxV5ApisApiSignRespData struct {
	AccessToken      string `json:"accessToken"`
	AccessExpiresAt  int64  `json:"accessExpiresAt"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresAt int64  `json:"refreshExpiresAt"`
}
