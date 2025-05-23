package qxTypes

type QxClientTestMsgReq struct {
	Msg string `json:"msg"`
}

type QxClientTestMsgResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type QxClientKeyExChangeReq struct {
	AccessKey string `json:"accessKey"`
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
	SessionId string `json:"sessionId"`
	PublicKey string `json:"publicKey"`
	ExpireAt  int64  `json:"expireAt"`
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

type QxClientApiTokenReq struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
	PublicKey    string `json:"publicKey"`
}

type QxClientApiTokenResp struct {
	Code      int32                     `json:"code"`
	Msg       string                    `json:"msg"`
	Path      string                    `json:"path"`
	RequestID string                    `json:"requestId"`
	Data      QxClientApiTokenRespModel `json:"data"`
}

type QxClientApiTokenRespModel struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
	TokenType   string `json:"tokenType"`
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
