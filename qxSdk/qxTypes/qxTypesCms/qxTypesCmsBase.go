package qxTypesCms

type CmsTestMsgReq struct {
	Msg string `json:"msg"`
}

type CmsEncryptResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type CmsTestMsgResp struct {
	Code      int32              `json:"code"`
	Msg       string             `json:"msg"`
	Path      string             `json:"path"`
	RequestID string             `json:"requestId"`
	Data      CmsTestMsgRespData `json:"data"`
}

type CmsTestMsgRespData struct {
	Msg string `json:"msg"`
}
