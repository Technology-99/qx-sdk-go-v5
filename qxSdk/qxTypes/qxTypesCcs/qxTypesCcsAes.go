package qxTypesCcs

type ModelAes struct {
	Key    string `json:"key"`
	AesKey string `json:"aesKey"`
	Iv     string `json:"iv"`
}

type GetAesFromKeyReq struct {
	Key string `json:"key"`
}

type GetAesFromKeyResp struct {
	Code      int32    `json:"code"`
	Msg       string   `json:"msg"`
	Path      string   `json:"path"`
	RequestID string   `json:"requestId"`
	Data      ModelAes `json:"data"`
}

type GetAesLatestReq struct {
}

type GetAesLatestResp struct {
	Code      int32    `json:"code"`
	Msg       string   `json:"msg"`
	Path      string   `json:"path"`
	RequestID string   `json:"requestId"`
	Data      ModelAes `json:"data"`
}
