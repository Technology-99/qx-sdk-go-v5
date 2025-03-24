package qxTypesKms

type KmsAkcCreateKeychainReq struct {
	CertType string `json:"certType"`
	Name     string `json:"name"`
}

type KmsAkcCreateKeychainResp struct {
	Code      int32                        `json:"code"`
	Msg       string                       `json:"msg"`
	Path      string                       `json:"path"`
	RequestID string                       `json:"requestId"`
	Data      KmsAkcCreateKeychainRespData `json:"data"`
}

type KmsAkcCreateKeychainRespData struct {
	Status     string `json:"status"`
	Name       string `json:"name"`
	SignMethod string `json:"signMethod"`
	CertType   string `json:"certType"`
	PublicKey  string `json:"publicKey"`
}

type KmsAkcSignReq struct {
	Name        string `json:"name"`
	SignContent string `json:"signContent"`
}

type KmsAkcSignResp struct {
	Code      int32              `json:"code"`
	Msg       string             `json:"msg"`
	Path      string             `json:"path"`
	RequestID string             `json:"requestId"`
	Data      KmsAkcSignRespData `json:"data"`
}

type KmsAkcSignRespData struct {
	Name string `json:"name"`
	Sign string `json:"sign"`
}

type KmsAkcVerifyReq struct {
	CertType string `json:"certType"`
	Name     string `json:"name"`
}

type KmsAkcVerifyResp struct {
	Code      int32                `json:"code"`
	Msg       string               `json:"msg"`
	Path      string               `json:"path"`
	RequestID string               `json:"requestId"`
	Data      KmsAkcVerifyRespData `json:"data"`
}

type KmsAkcVerifyRespData struct {
	Name   string `json:"name"`
	Verify bool   `json:"verify"`
}
