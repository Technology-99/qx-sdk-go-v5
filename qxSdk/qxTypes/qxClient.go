package qxTypes

type QxClientKeyExChangeReq struct {
	PublicKey string `json:"publicKey"`
}

type QxClientKeyExChangeResp struct {
	Code      int32                       `json:"code"`
	Msg       string                      `json:"msg"`
	Path      string                      `json:"path"`
	RequestID string                      `json:"requestId"`
	Data      QxClientKeyExChangeRespData `json:"data"`
}

type QxClientKeyExChangeRespData struct {
	PublicKey string `json:"publicKey"`
}

type QxClientApiDownPublicKeyResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type QxClientSignResultModel struct {
	AccessToken      string `json:"accessToken"`
	AccessExpiresAt  int64  `json:"accessExpiresAt"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresAt int64  `json:"refreshExpiresAt"`
}

type QxClientApiRefreshReq struct {
	AccessKey    string `json:"accessKey"`
	RefreshToken string `json:"refreshToken"`
}

type QxClientApiRefreshResp struct {
	Code      int32                   `json:"code"`
	Msg       string                  `json:"msg"`
	Path      string                  `json:"path"`
	RequestID string                  `json:"requestId"`
	Data      QxClientSignResultModel `json:"data"`
}

type QxClientApiSignReq struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

type QxClientApiSignResp struct {
	Code      int32                   `json:"code"`
	Msg       string                  `json:"msg"`
	Path      string                  `json:"path"`
	RequestID string                  `json:"requestId"`
	Data      QxClientSignResultModel `json:"data"`
}
type QxClientHealthzResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}
