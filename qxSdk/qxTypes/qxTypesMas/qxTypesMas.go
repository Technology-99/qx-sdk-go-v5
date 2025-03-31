package qxTypesMas

type SmsInitReq struct {
	Key     string `json:"key"`
	Service string `json:"service"`
	Type    string `json:"type,optional"`
	Zone    string `json:"zone,optional"`
	Mobile  string `json:"mobile,optional"`
}

type SmsInitResp struct {
	Code      int32           `json:"code"`
	Msg       string          `json:"msg"`
	Path      string          `json:"path"`
	RequestID string          `json:"requestId"`
	Data      SmsInitRespData `json:"data"`
}

type SmsInitRespData struct {
	Status string `json:"status"`
}

type SmsVerifyReq struct {
	Service    string `json:"service"`
	Type       string `json:"type,optional"`
	Zone       string `json:"zone,optional"`
	Mobile     string `json:"mobile,optional"`
	VerifyCode string `json:"verifyCode"`
}

type SmsVerifyResp struct {
	Code      int32             `json:"code"`
	Msg       string            `json:"msg"`
	Path      string            `json:"path"`
	RequestID string            `json:"requestId"`
	Data      SmsVerifyRespData `json:"data"`
}

type SmsVerifyRespData struct {
	Result bool `json:"result"`
}

type BehavioralVerificationInitReq struct {
	Key     string `json:"key"`
	Service string `json:"service"`
	Type    string `json:"type"`
}

type BehavioralVerificationInitResp struct {
	Code      int32                              `json:"code"`
	Msg       string                             `json:"msg"`
	Path      string                             `json:"path"`
	RequestID string                             `json:"requestId"`
	Data      BehavioralVerificationInitRespData `json:"data"`
}

type BehavioralVerificationInitRespData struct {
	Id      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	Answer  string `json:"answer,omitempty"`
	Img     string `json:"img,omitempty"`
}

type BehavioralVerificationVerifyReq struct {
	Id         string `json:"id"`
	Service    string `json:"service"`
	Type       string `json:"type"`
	VerifyCode string `json:"verifyCode"`
}

type BehavioralVerificationVerifyResp struct {
	Code      int32                                `json:"code"`
	Msg       string                               `json:"msg"`
	Path      string                               `json:"path"`
	RequestID string                               `json:"requestId"`
	Data      BehavioralVerificationVerifyRespData `json:"data"`
}

type BehavioralVerificationVerifyRespData struct {
	Result bool `json:"result"`
}
