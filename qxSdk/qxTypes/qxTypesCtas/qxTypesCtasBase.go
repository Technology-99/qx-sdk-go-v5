package qxTypesCtas

type CtasTestMsgReq struct {
	Msg string `json:"msg"`
}

type CtasTestMsgResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}
