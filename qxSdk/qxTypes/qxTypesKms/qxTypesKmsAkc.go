package qxTypesKms

// skc

type KmsSkcBatchDecryptReq struct {
	Data map[string]KmsSkcBatchDecryptReqItem `json:"data"`
}

type KmsSkcBatchDecryptReqItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcBatchDecryptResp struct {
	Code      int32                                 `json:"code"`
	Msg       string                                `json:"msg"`
	Path      string                                `json:"path"`
	RequestID string                                `json:"requestId"`
	Data      map[string]KmsSkcBatchDecryptRespItem `json:"data"`
}

type KmsSkcBatchDecryptRespItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
	Status   string `json:"status"`
}

type KmsSkcBatchEncryptReq struct {
	Data map[string]KmsSkcBatchEncryptReqItem `json:"data"`
}

type KmsSkcBatchEncryptReqItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcBatchEncryptResp struct {
	Code      int32                                 `json:"code"`
	Msg       string                                `json:"msg"`
	Path      string                                `json:"path"`
	RequestID string                                `json:"requestId"`
	Data      map[string]KmsSkcBatchEncryptRespItem `json:"data"`
}

type KmsSkcBatchEncryptRespItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
	Status   string `json:"status"`
}

type KmsSkcCompareItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcCompareReq struct {
	List []KmsSkcCompareItem `json:"list"`
}

type KmsSkcCompareResp struct {
	Code      int32                 `json:"code"`
	Msg       string                `json:"msg"`
	Path      string                `json:"path"`
	RequestID string                `json:"requestId"`
	Data      KmsSkcCompareRespData `json:"data"`
}

type KmsSkcCompareRespData struct {
	Status      string                      `json:"status"`
	CompareData bool                        `json:"compareData"`
	List        []KmsSkcCompareRespDataItem `json:"list"`
}

type KmsSkcCompareRespDataItem struct {
	Status   string `json:"status"`
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcCreateKeychainReq struct {
	Algorithm string `json:"algorithm"`
	Name      string `json:"name"`
}

type KmsSkcCreateKeychainResp struct {
	Code      int32                        `json:"code"`
	Msg       string                       `json:"msg"`
	Path      string                       `json:"path"`
	RequestID string                       `json:"requestId"`
	Data      KmsSkcCreateKeychainRespData `json:"data"`
}

type KmsSkcCreateKeychainRespData struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

type KmsSkcDecryptReq struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcDecryptResp struct {
	Code      int32                 `json:"code"`
	Msg       string                `json:"msg"`
	Path      string                `json:"path"`
	RequestID string                `json:"requestId"`
	Data      KmsSkcDecryptRespData `json:"data"`
}

type KmsSkcDecryptRespData struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcEncryptReq struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcEncryptResp struct {
	Code      int32                 `json:"code"`
	Msg       string                `json:"msg"`
	Path      string                `json:"path"`
	RequestID string                `json:"requestId"`
	Data      KmsSkcEncryptRespData `json:"data"`
}

type KmsSkcEncryptRespData struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

// akc
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
