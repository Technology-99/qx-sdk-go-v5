package qxTypesCcs

type CcsTestMsgReq struct {
	Key string `json:"key"`
	Msg string `json:"msg"`
}

type CcsEncryptResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type CcsTestMsgResp struct {
	Code      int32              `json:"code"`
	Msg       string             `json:"msg"`
	Path      string             `json:"path"`
	RequestID string             `json:"requestId"`
	Data      CcsTestMsgRespData `json:"data"`
}

type CcsTestMsgRespData struct {
	Msg string `json:"msg"`
}
