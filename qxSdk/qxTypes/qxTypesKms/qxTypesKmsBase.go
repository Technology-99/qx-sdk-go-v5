package qxTypesKms

type KmsTestMsgReq struct {
	Msg string `json:"msg"`
}

type KmsEncryptResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}
type KmsTestMsgResp struct {
	Code      int32              `json:"code"`
	Msg       string             `json:"msg"`
	Path      string             `json:"path"`
	RequestID string             `json:"requestId"`
	Data      KmsTestMsgRespData `json:"data"`
}

type KmsTestMsgRespData struct {
	Msg string `json:"msg"`
}
